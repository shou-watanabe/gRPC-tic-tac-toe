package entity

import (
	"sync"
)

type TicTacToe struct {
	sync.RWMutex
	started  bool
	finished bool
	me       *Player
	room     *Room
	game     *Game
}
