package game

import "fmt"

// import "fmt"

type Board struct {
	Cells [][]Symbol
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
