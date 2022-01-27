package handler

import (
	"fmt"
	"sync"

	"gRPC-tic-tac-toe/build"
	"gRPC-tic-tac-toe/domain/entity"
	"gRPC-tic-tac-toe/gen/pb"
	repository "gRPC-tic-tac-toe/infra"
	"gRPC-tic-tac-toe/usecase"
)

type gameHandler struct {
	sync.RWMutex
	gameUsecase usecase.GameUsecase
	games       map[int32]*entity.Game                // ゲーム情報（盤面など）を格納する
	client      map[int32][]pb.GameService_PlayServer // 状態変更時にクライアントにストリーミングを返すために格納する
}

func NewGameHandler(gu usecase.GameUsecase) *gameHandler {
	return &gameHandler{
		games:       make(map[int32]*entity.Game),
		client:      make(map[int32][]pb.GameService_PlayServer),
		gameUsecase: gu,
	}
}

func (h *gameHandler) Play(stream pb.GameService_PlayServer) error {
	for {
		//クライアントからリクエストを受信したらreqにリクエストが代入されます
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		roomID := req.GetRoomId()
		player := build.Player(req.GetPlayer())

		//oneofで複数の型のリクエストがくるのでswtich文で処理します
		switch req.GetAction().(type) {
		case *pb.PlayRequest_Start:
			//ゲーム開始リクエスト
			err := h.start(stream, roomID, player)
			if err != nil {
				return err
			}
		case *pb.PlayRequest_Move:
			//石を置いた時のリクエスト
			action := req.GetMove()
			x := action.GetMove().GetX()
			y := action.GetMove().GetY()
			err := h.move(roomID, x, y, player)
			if err != nil {
				return err
			}
		}
	}
}

func (h *gameHandler) start(stream pb.GameService_PlayServer, roomID int32, me *entity.Player) error {
	h.Lock()
	defer h.Unlock()

	//ゲーム情報がなければ作成する
	g := h.games[roomID]
	if g == nil {
		g = repository.NewGame(entity.None)
		h.games[roomID] = g
		h.client[roomID] = make([]pb.GameService_PlayServer, 0, 2)
	}

	//自分のクライアントを格納する
	h.client[roomID] = append(h.client[roomID], stream)

	if len(h.client[roomID]) == 2 {
		// 二人揃ったので開始する
		for _, s := range h.client[roomID] {
			// クライアントにゲーム開始を通知する
			err := s.Send(&pb.PlayResponse{
				Event: &pb.PlayResponse_Ready{
					Ready: &pb.PlayResponse_ReadyEvent{},
				},
			})
			if err != nil {
				return err
			}
		}
		fmt.Printf("game has started room_id=%v\n", roomID)
	} else {
		//まだ揃ってないので待機中であることをクライアントに通知する
		err := stream.Send(&pb.PlayResponse{
			Event: &pb.PlayResponse_Waiting{
				Waiting: &pb.PlayResponse_WaitingEvent{},
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *gameHandler) move(roomID int32, x int32, y int32, p *entity.Player) error {
	h.Lock()
	defer h.Unlock()

	g := h.games[roomID]

	fmt.Println("ここでエラー？")
	finished, err := h.gameUsecase.Move(x, y, p.Symbol, g)
	if err != nil {
		return err
	}

	for _, s := range h.client[roomID] {
		// 手が打たれたことをクライアントに通知する
		err := s.Send(&pb.PlayResponse{
			Event: &pb.PlayResponse_Move{
				Move: &pb.PlayResponse_MoveEvent{
					Player: build.PBPlayer(p),
					Move: &pb.Move{
						X: x,
						Y: y,
					},
					Board: build.PBBoard(g.Board),
				},
			},
		})
		if err != nil {
			return err
		}

		if finished {
			// ゲーム終了通知する
			err := s.Send(
				&pb.PlayResponse{
					Event: &pb.PlayResponse_Finished{
						Finished: &pb.PlayResponse_FinishedEvent{
							Winner: build.PBSymbol(h.gameUsecase.Winner(g)),
							Board:  build.PBBoard(g.Board),
						},
					},
				},
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
