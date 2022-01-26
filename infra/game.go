package repository

import (
	"fmt"
	"gRPC-tic-tac-toe/domain/entity"
	"gRPC-tic-tac-toe/domain/repository"
)

type gameRepository struct {
	sr repository.SymbolRepository
	br repository.BoardRepository
}

func NewGameRepository(sr repository.SymbolRepository, br repository.BoardRepository) repository.GameRepository {
	return &gameRepository{sr: sr, br: br}
}

func NewGame(me entity.Symbol) *entity.Game {
	return &entity.Game{
		// Board: NewBoard(),
		Me: me,
	}
}

// 手を打ちます。その後盤面を出力します。
// 返り値として、ゲームが終了したかどうかを返します。
func (gr *gameRepository) Move(x int32, y int32, s entity.Symbol, g *entity.Game) (bool, error) {
	if g.Finished {
		return true, nil
	}
	err := gr.br.PutStone(x-1, y-1, s, g.Board)
	if err != nil {
		return false, err
	}
	gr.Display(g.Me, g)
	if gr.IsGameOver(g) != entity.NoWin {
		fmt.Println("finished")
		g.Finished = true
		return true, nil
	}

	return false, nil
}

// ゲームが終了したかを判定します
// 黒と白双方に置ける場所がなければ終了とします
// usecase?
func (gr *gameRepository) IsGameOver(g *entity.Game) entity.Winner {
	if !gr.br.IsAvailableEmpty(g.Board) {
		return entity.Draw
	}

	if gr.br.IsAvailableLine(entity.Circle, g.Board) {
		return entity.CircleWin
	}

	if gr.br.IsAvailableLine(entity.Cross, g.Board) {
		return entity.CrossWin
	}

	return entity.NoWin
}

// 勝者の色を返します。引き分けの場合はNoneを返します
// usecase?
func (gr *gameRepository) Winner(g *entity.Game) entity.Symbol {
	if gr.br.IsAvailableLine(entity.Circle, g.Board) {
		return entity.Circle
	}

	if gr.br.IsAvailableLine(entity.Cross, g.Board) {
		return entity.Cross
	}

	return entity.None
}

// 盤面を出力します
// usecase?
func (gr *gameRepository) Display(me entity.Symbol, g *entity.Game) {
	fmt.Println("")
	if me != entity.None {
		fmt.Printf("You: %v\n", gr.sr.SymbolToStr(me))
	}

	fmt.Print(" ｜")
	rs := []rune("ABC")
	for i, r := range rs {
		fmt.Printf("%v", string(r))
		if i < len(rs)-1 {
			fmt.Print("｜")
		}
	}
	fmt.Print("｜")
	fmt.Print("\n")
	fmt.Println("ーーーーーー")

	for j := 0; j < 3; j++ {
		fmt.Printf("%d", j+1)
		fmt.Print("｜")
		for i := 0; i < 3; i++ {
			fmt.Print(gr.sr.SymbolToStr(g.Board.Cells[i][j]))
			fmt.Print("｜")
		}
		fmt.Print("\n")
	}

	fmt.Println("ーーーーーー")

	fmt.Print("\n")
}
