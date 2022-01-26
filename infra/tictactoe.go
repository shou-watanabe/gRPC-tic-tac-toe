package repository

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"gRPC-tic-tac-toe/build"
	"gRPC-tic-tac-toe/domain/entity"
	"gRPC-tic-tac-toe/domain/repository"
	"gRPC-tic-tac-toe/gen/pb"
)

type ticTacToeRepository struct {
	gr repository.GameRepository
}

func NewTicTacToeRepository(gr repository.GameRepository) repository.TicTacToeRepository {
	return &ticTacToeRepository{gr: gr}
}

func NewTicTacToe() *entity.TicTacToe {
	return &entity.TicTacToe{}
}

func (tr ticTacToeRepository) Run(t *entity.TicTacToe, g *entity.Game) int {
	if err := tr.PreRun(t, g); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func (tr ticTacToeRepository) PreRun(t *entity.TicTacToe, g *entity.Game) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "Failed to connect to grpc server")
	}
	defer conn.Close()

	// マッチング問い合わせ
	err = tr.Matching(ctx, pb.NewMatchingServiceClient(conn), t)
	if err != nil {
		return err
	}

	// マッチングできたので盤面生成
	t.Game = NewGame(t.Me.Symbol)

	// 双方向ストリーミングでゲーム処理
	return tr.ExecPlay(ctx, pb.NewGameServiceClient(conn), t, g)
}

func (tr ticTacToeRepository) Matching(ctx context.Context, cli pb.MatchingServiceClient, t *entity.TicTacToe) error {
	// マッチングリクエスト
	stream, err := cli.JoinRoom(ctx, &pb.JoinRoomRequest{})
	if err != nil {
		return err
	}
	defer stream.CloseSend()

	fmt.Println("Requested matching...")

	// ストリーミングでレスポンスを受け取る
	for {
		resp, err := stream.Recv()
		if err != nil {
			return err
		}

		if resp.GetStatus() == pb.JoinRoomResponse_MATCHED {
			// マッチング成立
			t.Room = build.Room(resp.GetRoom())
			t.Me = build.Player(resp.GetMe())
			fmt.Printf("Matched room_id=%d\n", resp.GetRoom().GetId())
			return nil
		} else if resp.GetStatus() == pb.JoinRoomResponse_WAITTING {
			// 待機中
			fmt.Println("Waiting mathing...")
		}
	}
}

func (tr ticTacToeRepository) ExecPlay(ctx context.Context, cli pb.GameServiceClient, t *entity.TicTacToe, g *entity.Game) error {
	c, cancel := context.WithCancel(ctx)
	defer cancel()

	// 双方向ストリーミングを開始する
	stream, err := cli.Play(c)
	if err != nil {
		return err
	}
	defer stream.CloseSend()

	go func() {
		err := tr.Send(c, stream, t, g)
		if err != nil {
			cancel()
		}
	}()

	err = tr.Receive(c, stream, t, g)
	if err != nil {
		cancel()
		return err
	}

	return nil
}

func (tr ticTacToeRepository) Play(t *entity.TicTacToe, g *entity.Game) (bool, error) {
	tr.gr.Display(1, g)
	fmt.Print("Input Your Move (ex. A-1):")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	// 入力された手を解析する
	text := stdin.Text()
	x, y, err := parseInput(text)
	if err != nil {
		return false, err
	}
	isGameOver, err := tr.gr.Move(x-1, y-1, t.Me.Symbol, g)
	if err != nil {
		return false, err
	}

	tr.gr.Display(1, g)

	return isGameOver, nil
}

// `A-2`の形式で入力された手を (x, y)=(1,2) の形式に変換する
func parseInput(txt string) (int32, int32, error) {
	ss := strings.Split(txt, "-")
	if len(ss) != 2 {
		return 0, 0, fmt.Errorf("入力が不正です。例: A-1")
	}

	xs := ss[0]
	xrs := []rune(strings.ToUpper(xs))
	x := int32(xrs[0]-rune('A')) + 1

	if x < 1 || 8 < x {
		return 0, 0, fmt.Errorf("入力が不正です。例: A-1")
	}

	ys := ss[1]
	y, err := strconv.ParseInt(ys, 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("入力が不正です。例: A-1")
	}
	if y < 1 || 8 < y {
		return 0, 0, fmt.Errorf("入力が不正です。例: A-1")
	}

	return x, int32(y), nil
}

func (tr ticTacToeRepository) Receive(ctx context.Context, stream pb.GameService_PlayClient, t *entity.TicTacToe, g *entity.Game) error {
	for {
		// サーバーからのストリーミングを受け取る
		res, err := stream.Recv()
		if err != nil {
			return err
		}

		t.Lock()
		switch res.GetEvent().(type) {
		case *pb.PlayResponse_Waiting:
			// 開始待機中
		case *pb.PlayResponse_Ready:
			// 開始
			t.Started = true
			tr.gr.Display(t.Me.Symbol, g)
		case *pb.PlayResponse_Move:
			// 手を打たれた
			color := build.Symbol(res.GetMove().GetPlayer().GetSymbol())
			if color != t.Me.Symbol {
				move := res.GetMove().GetMove()
				// クライアント側のゲーム情報に反映させる
				tr.gr.Move(move.GetX(), move.GetY(), color, g)
				fmt.Print("Input Your Move (ex. A-1):")
			}
		case *pb.PlayResponse_Finished:
			// ゲームが終了した
			t.Finished = true

			// 勝敗を表示する
			winner := build.Symbol(res.GetFinished().Winner)
			fmt.Println("")
			if winner == entity.None {
				fmt.Println("Draw!")
			} else if winner == t.Me.Symbol {
				fmt.Println("You Win!")
			} else {
				fmt.Println("You Lose!")
			}

			// ループを終了する
			t.Unlock()
			return nil
		}
		t.Unlock()

		select {
		case <-ctx.Done():
			// キャンセルされたので終了する
			return nil
		default:
		}
	}
}

func (tr ticTacToeRepository) Send(ctx context.Context, stream pb.GameService_PlayClient, t *entity.TicTacToe, g *entity.Game) error {
	for {
		t.RLock()

		if t.Finished {
			// recieve側で終了されたので、send側も終了する
			t.RUnlock()
			return nil
		} else if !t.Started {
			// 未開始なので、開始リクエストを送る
			err := stream.Send(&pb.PlayRequest{
				RoomId: t.Room.ID,
				Player: build.PBPlayer(t.Me),
				Action: &pb.PlayRequest_Start{
					Start: &pb.PlayRequest_StartAction{},
				},
			})
			t.RUnlock()
			if err != nil {
				return err
			}

			for {
				// 相手が開始するまで待機する
				t.RLock()
				if t.Started {
					// 開始をrecieveした
					t.RUnlock()
					fmt.Println("READY GO!")
					break
				}
				t.RUnlock()
				fmt.Println("Waiting until opponent player ready")
				time.Sleep(1 * time.Second)
			}
		} else {
			// 対戦中

			t.RUnlock()
			// 手の入力を待機する
			fmt.Print("Input Your Move (ex. A-1):")
			stdin := bufio.NewScanner(os.Stdin)
			stdin.Scan()

			// 入力された手を解析する
			text := stdin.Text()
			x, y, err := parseInput(text)
			if err != nil {
				fmt.Println(err)
				continue
			}

			// 手を打つ
			t.Lock()
			_, err = tr.gr.Move(x, y, t.Me.Symbol, g)
			t.Unlock()
			if err != nil {
				fmt.Println(err)
				continue
			}

			go func() {
				// サーバーに手を送る
				err = stream.Send(&pb.PlayRequest{
					RoomId: t.Room.ID,
					Player: build.PBPlayer(t.Me),
					Action: &pb.PlayRequest_Move{
						Move: &pb.PlayRequest_MoveAction{
							Move: &pb.Move{
								X: x,
								Y: y,
							},
						},
					},
				})
				if err != nil {
					fmt.Println(err)
				}
			}()

			// 一度手を打ったら5秒間待機する
			ch := make(chan int)
			go func(ch chan int) {
				fmt.Println("")
				for i := 0; i < 5; i++ {
					fmt.Printf("freezing in %d second.\n", (5 - i))
					time.Sleep(1 * time.Second)
				}
				fmt.Println("")
				ch <- 0
			}(ch)
			<-ch
		}

		select {
		case <-ctx.Done():
			// キャンセルされたので終了する
			return nil
		default:
		}
	}
}
