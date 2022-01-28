package repository

import (
	"fmt"

	"gRPC-tic-tac-toe/config"
	"gRPC-tic-tac-toe/domain/entity"
	"gRPC-tic-tac-toe/domain/repository"
)

type boardRepository struct {
	sr repository.SymbolRepository
}

func NewBoardRepository(sr repository.SymbolRepository) repository.BoardRepository {
	return &boardRepository{sr: sr}
}

// 盤面を作成します
func NewBoard() *entity.Board {
	// 3x3のマスの盤面を二次元配列で作ります
	b := &entity.Board{
		Cells: make([][]entity.Symbol, 3),
	}
	for i := 0; i < 3; i++ {
		b.Cells[i] = make([]entity.Symbol, 3)
	}

	return b
}

// 石を置きます
func (br *boardRepository) PutStone(x int32, y int32, s entity.Symbol, b *entity.Board) error {
	// そのマスに石を置けるかチェックします
	if !br.CanPutStone(x, y, b) {
		return fmt.Errorf("can not put stone x=%v, y=%v symbol=%v", x, y, br.sr.SymbolToStr(s))
	}

	// マスに石を置きます
	b.Cells[x][y] = s

	return nil
}

// マスに石を置けるか判定します
func (br *boardRepository) CanPutStone(x int32, y int32, b *entity.Board) bool {
	// すでに石がある場合は石を置けません
	return b.Cells[x][y] == entity.Empty
}

// 空きマスがあるか判定
func (br *boardRepository) IsAvailableEmpty(b *entity.Board) bool {
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; j++ {
			// 空きマスが一つでもあればtrue
			if b.Cells[i][j] == entity.Empty {
				return true
			}
		}
	}

	// 空きマスが一つもなければfalse
	return false
}

// ラインができているか判定
func (br *boardRepository) IsAvailableLine(s entity.Symbol, b *entity.Board) bool {
	for _, line := range config.Lines {
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
