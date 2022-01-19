package game

import "fmt"

type Board struct {
	Cells [][]Symbol
}

// 勝ち手
var lines = [][][]int{
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
	{{0, 0}, {1, 1}, {2, 2}},
	{{0, 2}, {1, 1}, {2, 0}},
}

// 盤面を作成します
func NewBoard() *Board {
	// 3x3のマスの盤面を二次元配列で作ります
	b := &Board{
		Cells: make([][]Symbol, 3),
	}
	for i := 0; i < 3; i++ {
		b.Cells[i] = make([]Symbol, 3)
	}

	return b
}

// 石を置きます
func (b *Board) PutStone(x int32, y int32, s Symbol) error {
	// そのマスに石を置けるかチェックします
	if !b.CanPutStone(x, y) {
		return fmt.Errorf("can not put stone x=%v, y=%v symbol=%v", x, y, SymbolToStr(s))
	}

	// マスに石を置きます
	b.Cells[x][y] = s

	return nil
}

// マスに石を置けるか判定します
func (b *Board) CanPutStone(x int32, y int32) bool {
	// すでに石がある場合は石を置けません
	return b.Cells[x][y] == Empty
}

// 空きマスがあるか判定
func (b *Board) IsAvailableEmpty() bool {
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; j++ {
			// 空きマスが一つでもあればtrue
			if b.Cells[i][j] == Empty {
				return true
			}
		}
	}

	// 空きマスが一つもなければfalse
	return false
}

// ラインができているか判定
func (b *Board) IsAvailableLine(s Symbol) bool {
	for _, line := range lines {
		for i, cell := range line {
			// 勝ち手に合わない場合にbreak
			if b.Cells[cell[0]][cell[1]] != s {
				break
			}

			// 最後まで勝ち手に合えばtrueを返す
			if i == len(line)-1 {
				return true
			}
		}
	}

	// 一つも勝ち手に当てはまらなければfalseを返す
	return false
}
