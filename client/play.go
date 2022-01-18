package client

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gRPC-tic-tac-toe/game"
)

type TicTacToe struct {
	// sync.RWMutex
	started  bool
	finished bool
	me       *game.Player
	// room     *game.Room
	game *game.Game
}

func NewTicTacToe(g *game.Game, me *game.Player) *TicTacToe {
	return &TicTacToe{game: g, me: me}
}

func (t *TicTacToe) Play() (bool, error) {
	t.game.Display(1)
	fmt.Print("Input Your Move (ex. A-1):")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	// 入力された手を解析する
	text := stdin.Text()
	x, y, err := parseInput(text)
	if err != nil {
		return false, err
	}
	isGameOver, err := t.game.Move(x-1, y-1, t.me.Symbol)
	if err != nil {
		return false, err
	}

	t.game.Display(1)

	return isGameOver, nil
}

// `A-2`の形式で入力された手を (x, y)=(1,2) の形式に変換する
func parseInput(txt string) (int32, int32, error) {
	ss := strings.Split(txt, "-")
	if len(ss) != 2 {
		return 0, 0, fmt.Errorf("入力が不正です。例：A-1")
	}

	xs := ss[0]
	xrs := []rune(strings.ToUpper(xs))
	x := int32(xrs[0]-rune('A')) + 1

	if x < 1 || 8 < x {
		return 0, 0, fmt.Errorf("入力が不正です。例：A-1")
	}

	ys := ss[1]
	y, err := strconv.ParseInt(ys, 10, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("入力が不正です。例：A-1")
	}
	if y < 1 || 8 < y {
		return 0, 0, fmt.Errorf("入力が不正です。例：A-1")
	}

	return x, int32(y), nil
}
