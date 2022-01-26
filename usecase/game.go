package usecase

import (
	"gRPC-tic-tac-toe/domain/entity"
	"gRPC-tic-tac-toe/domain/repository"
)

type GameUsecase interface {
	Move(x int32, y int32, s entity.Symbol, g *entity.Game) (bool, error)
	Winner(g *entity.Game) entity.Symbol
}

type gameUsecase struct {
	gameRepository repository.GameRepository
}

func NewGameUsecase(gr repository.GameRepository) GameUsecase {
	gameUsecase := gameUsecase{gameRepository: gr}
	return &gameUsecase
}

func (gu gameUsecase) Move(x int32, y int32, s entity.Symbol, g *entity.Game) (bool, error) {
	return gu.gameRepository.Move(x, y, s, g)
}

func (gu gameUsecase) Winner(g *entity.Game) entity.Symbol {
	return gu.gameRepository.Winner(g)
}
