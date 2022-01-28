package build

import (
	"fmt"

	"gRPC-tic-tac-toe/domain/entity"
	"gRPC-tic-tac-toe/gen/pb"
)

func Room(r *pb.Room) *entity.Room {
	return &entity.Room{
		ID:    r.GetId(),
		Host:  Player(r.GetHost()),
		Guest: Player(r.GetGuest()),
	}
}

func Player(p *pb.Player) *entity.Player {
	return &entity.Player{
		ID:     p.GetId(),
		Symbol: Symbol(p.GetSymbol()),
	}
}

func Symbol(c pb.Symbol) entity.Symbol {
	switch c {
	case pb.Symbol_CIRCLE:
		return entity.Circle
	case pb.Symbol_CROSS:
		return entity.Cross
	case pb.Symbol_EMPTY:
		return entity.Empty
	}

	panic(fmt.Sprintf("unknwon symbol=%v", c))
}
