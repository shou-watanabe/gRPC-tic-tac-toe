package build

import (
	"gRPC-tic-tac-toe/domain/entity"
	"gRPC-tic-tac-toe/gen/pb"
)

func PBRoom(r *entity.Room) *pb.Room {
	return &pb.Room{
		Id:    r.ID,
		Host:  PBPlayer(r.Host),
		Guest: PBPlayer(r.Guest),
	}
}

func PBPlayer(p *entity.Player) *pb.Player {
	if p == nil {
		return nil
	}
	return &pb.Player{
		Id:     p.ID,
		Symbol: PBSymbol(p.Symbol),
	}
}

func PBSymbol(c entity.Symbol) pb.Symbol {
	switch c {
	case entity.Circle:
		return pb.Symbol_CIRCLE
	case entity.Cross:
		return pb.Symbol_CROSS
	case entity.Empty:
		return pb.Symbol_EMPTY
	}

	return pb.Symbol_UNKNOWN
}

func PBBoard(b *entity.Board) *pb.Board {
	pbCols := make([]*pb.Board_Sym, 0, 10)

	for _, col := range b.Cells {
		pbCells := make([]pb.Symbol, 0, 10)
		for _, c := range col {
			pbCells = append(pbCells, PBSymbol(c))
		}
		pbCols = append(pbCols, &pb.Board_Sym{
			Cells: pbCells,
		})
	}

	return &pb.Board{
		Cols: pbCols,
	}
}
