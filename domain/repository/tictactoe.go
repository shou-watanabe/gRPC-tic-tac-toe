package repository

import (
	"context"

	"gRPC-tic-tac-toe/domain/entity"
	"gRPC-tic-tac-toe/gen/pb"
)

type TicTacToeRepository interface {
	Run(t *entity.TicTacToe, g *entity.Game) int
	PreRun(t *entity.TicTacToe, g *entity.Game) error
	Matching(ctx context.Context, cli pb.MatchingServiceClient, t *entity.TicTacToe) error
	ExecPlay(ctx context.Context, cli pb.GameServiceClient, t *entity.TicTacToe, g *entity.Game) error
	Play(t *entity.TicTacToe, g *entity.Game) (bool, error)
	Receive(ctx context.Context, stream pb.GameService_PlayClient, t *entity.TicTacToe, g *entity.Game) error
	Send(ctx context.Context, stream pb.GameService_PlayClient, t *entity.TicTacToe, g *entity.Game) error
}
