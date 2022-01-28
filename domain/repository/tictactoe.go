package repository

import (
	"context"

	"gRPC-tic-tac-toe/domain/entity"
	"gRPC-tic-tac-toe/gen/pb"
)

type TicTacToeRepository interface {
	Run(t *entity.TicTacToe) int
	PreRun(t *entity.TicTacToe) error
	Matching(ctx context.Context, cli pb.MatchingServiceClient, t *entity.TicTacToe) error
	ExecPlay(ctx context.Context, cli pb.GameServiceClient, t *entity.TicTacToe) error
	Play(t *entity.TicTacToe) (bool, error)
	Receive(ctx context.Context, stream pb.GameService_PlayClient, t *entity.TicTacToe) error
	Send(ctx context.Context, stream pb.GameService_PlayClient, t *entity.TicTacToe) error
}
