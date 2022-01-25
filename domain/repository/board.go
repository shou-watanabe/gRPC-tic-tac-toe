package repository

import "gRPC-tic-tac-toe/domain/entity"

type BoardRepository interface {
	PutStone(x int32, y int32, s entity.Symbol, b *entity.Board) error
	CanPutStone(x int32, y int32, b *entity.Board) bool
	IsAvailableEmpty(b *entity.Board) bool
	IsAvailableLine(s entity.Symbol, b *entity.Board) bool
}
