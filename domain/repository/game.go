package repository

import (
	"gRPC-tic-tac-toe/domain/entity"
)

type GameRepository interface {
	Move(x int32, y int32, c entity.Symbol, g entity.Game, b entity.Board) (bool, error)
	IsGameOver(g entity.Game, b entity.Board) entity.Winner
	Winner(g entity.Game, b entity.Board) entity.Symbol
	Display(me entity.Symbol, g entity.Game)
}
