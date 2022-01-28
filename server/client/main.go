package main

import (
	"os"

	repository "gRPC-tic-tac-toe/infra"
)

func main() {
	sr := repository.NewSymbolRepository()
	br := repository.NewBoardRepository(sr)
	gr := repository.NewGameRepository(sr, br)
	tr := repository.NewTicTacToeRepository(gr)
	os.Exit(tr.Run(repository.NewTicTacToe()))
}
