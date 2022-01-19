package main

import (
	"os"

	"gRPC-tic-tac-toe/client"
)

func main() {
	os.Exit(client.NewTicTacToe().Run())
}
