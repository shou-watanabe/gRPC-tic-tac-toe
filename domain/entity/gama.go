package entity

type Game struct {
	Board    *Board
	Finished bool
	Me       Symbol
}
