package repository

import "gRPC-tic-tac-toe/domain/entity"

type SymbolRepository interface {
	SymbolToStr(s entity.Symbol) string
	OpponentSymbol(me entity.Symbol) entity.Symbol
}
