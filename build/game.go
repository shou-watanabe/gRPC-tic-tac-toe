package build

import (
	"fmt"

	"gRPC-tic-tac-toe/game"
	"gRPC-tic-tac-toe/gen/pb"
)

func Room(r *pb.Room) *game.Room {
	return &game.Room{
		ID:    r.GetId(),
		Host:  Player(r.GetHost()),
		Guest: Player(r.GetGuest()),
	}
}

func Player(p *pb.Player) *game.Player {
	return &game.Player{
		ID:     p.GetId(),
		Symbol: Symbol(p.GetSymbol()),
	}
}

func Symbol(c pb.Symbol) game.Symbol {
	switch c {
	case pb.Symbol_CIRCLE:
		return game.Circle
	case pb.Symbol_CROSS:
		return game.Cross
	case pb.Symbol_EMPTY:
		return game.Empty
	}

	panic(fmt.Sprintf("unknwon symbol=%v", c))
}
