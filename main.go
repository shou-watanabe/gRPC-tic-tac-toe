package main

import "gRPC-tic-tac-toe/game"

func main() {
	g := game.NewGame(1)
	g.Display(1)
}
