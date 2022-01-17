package game

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

// // 石を置きます
// func (b *Board) PutStone(x int32, y int32, c Color) error {
// 	// そのマスに石を置けるかチェックします
// 	if !b.CanPutStone(x, y, c) {
// 		return fmt.Errorf("Can not put stone x=%v, y=%v color=%v", x, y, ColorToStr(c))
// 	}

// 	// マスに石を置きます
// 	b.Cells[x][y] = c

// 	// 置いた石の縦/横/斜めの各方向でひっくり返すことのできる石を全てひっくり返します
// 	for dx := -1; dx <= 1; dx++ {
// 		for dy := -1; dy <= 1; dy++ {
// 			if dx == 0 && dy == 0 {
// 				continue
// 			}
// 			if b.CountTurnableStonesByDirection(x, y, c, int32(dx), int32(dy)) > 0 {
// 				b.TurnStonesByDirection(x, y, c, int32(dx), int32(dy))
// 			}
// 		}
// 	}

// 	return nil
// }

// // マスに石を置けるか判定します
// func (b *Board) CanPutStone(x int32, y int32, c Color) bool {
// 	// すでに石がある場合は石を置けません
// 	if b.Cells[x][y] != Empty {
// 		return false
// 	}

// 	// 縦/横/斜めの各方向をチェックします
// 	for dx := -1; dx <= 1; dx++ {
// 		for dy := -1; dy <= 1; dy++ {
// 			if dx == 0 && dy == 0 {
// 				continue
// 			}

// 			// ひっくり返すことのできる石がひとつでもあれば、石を置けます
// 			if b.CountTurnableStonesByDirection(x, y, c, int32(dx), int32(dy)) > 0 {
// 				return true
// 			}
// 		}
// 	}

// 	// ひとつもひっくり返すことできる石がなければ、石を置けません
// 	return false
// }
