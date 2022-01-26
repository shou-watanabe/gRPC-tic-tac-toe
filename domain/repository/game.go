package repository

import (
	"gRPC-tic-tac-toe/domain/entity"
)

type GameRepository interface {
	Move(x int32, y int32, s entity.Symbol, g *entity.Game) (bool, error)
	IsGameOver(g *entity.Game) entity.Winner
	Winner(g *entity.Game) entity.Symbol
	Display(me entity.Symbol, g *entity.Game)
}
