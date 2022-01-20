package repository

import (
	"gRPC-tic-tac-toe/domain/entity"
	"gRPC-tic-tac-toe/domain/repository"
)

type symbolRepository struct{}

func NewSymbolRepository() repository.SymbolRepository {
	return &symbolRepository{}
}

// 色を文字列に変換します
func (sr *symbolRepository) SymbolToStr(se entity.Symbol) string {
	switch se {
	case entity.Circle:
		return "○"
	case entity.Cross:
		return "×"
	case entity.Empty:
		return " "
	}

	return ""
}

// 対戦相手の色を取得します
func (sr *symbolRepository) OpponentSymbol(me entity.Symbol) entity.Symbol {
	switch me {
	case entity.Circle:
		return entity.Cross
	case entity.Cross:
		return entity.Circle
	}

	panic("invalid state")
}
