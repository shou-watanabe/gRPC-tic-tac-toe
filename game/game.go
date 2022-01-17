package game

import "fmt"

type Game struct {
	Board    *Board
	started  bool
	finished bool
	me       Symbol
}

type Winner int

const (
	Draw      Winner = iota // 誰も打ってない
	CircleWin               // マル
	CrossWin                // バツ
	NoWin                   // なんでもない
)

func NewGame(me Symbol) *Game {
	return &Game{
		Board: NewBoard(),
		me:    me,
	}
}

// // 手を打ちます。その後盤面を出力します。
// // 返り値として、ゲームが終了したかどうかを返します。
// func (g *Game) Move(x int32, y int32, c Symbol) (bool, error) {
// 	if g.finished {
// 		return true, nil
// 	}
// 	err := g.Board.PutStone(x, y, c)
// 	if err != nil {
// 		return false, err
// 	}
// 	g.Display(g.me)
// 	if g.IsGameOver() {
// 		fmt.Println("finished")
// 		g.finished = true
// 		return true, nil
// 	}

// 	return false, nil
// }

// ゲームが終了したかを判定します
// 黒と白双方に置ける場所がなければ終了とします
func (g *Game) IsGameOver() Winner {
	if !g.Board.IsAvailableEmpty() {
		return Draw
	}

	if g.Board.IsAvailableLine(Circle) {
		return CircleWin
	}

	if g.Board.IsAvailableLine(Cross) {
		return CrossWin
	}

	return NoWin
}

// //　勝者の色を返します。引き分けの場合はNoneを返します
// func (g *Game) Winner() Symbol {
// 	black := g.Board.Score(Black)
// 	white := g.Board.Score(White)
// 	if black == white {
// 		return None
// 	} else if black > white {
// 		return Black
// 	}
// 	return White
// }

// 盤面を出力します
func (g *Game) Display(me Symbol) {
	fmt.Println("")
	if me != None {
		fmt.Printf("You: %v\n", SymbolToStr(me))
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
			fmt.Print(SymbolToStr(g.Board.Cells[i][j]))
			fmt.Print("｜")
		}
		fmt.Print("\n")
	}

	fmt.Println("ーーーーーー")

	fmt.Print("\n")
}
