package game

import "fmt"

type Game struct {
	Board    *Board
	started  bool
	finished bool
	me       Symbol
}

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

// // ゲームが終了したかを判定します
// // 黒と白双方に置ける場所がなければ終了とします
// func (g *Game) IsGameOver() bool {
// 	if g.Board.AvailableCellCount(Black) > 0 {
// 		return false
// 	}

// 	if g.Board.AvailableCellCount(White) > 0 {
// 		return false
// 	}

// 	return true
// }

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

	// fmt.Printf("Score: BLACK=%d, WHITE=%d REST=%d\n", g.Board.Score(Black), g.Board.Score(White), g.Board.Rest())

	fmt.Print("\n")
}
