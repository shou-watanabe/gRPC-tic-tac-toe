package entity

import (
	"sync"
)

type TicTacToe struct {
	sync.RWMutex
	Started  bool
	Finished bool
	Me       *Player
	Room     *Room
	Game     *Game
}
