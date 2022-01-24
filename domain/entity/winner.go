package entity

type Winner int

const (
	Draw      Winner = iota // 誰も打ってない
	CircleWin               // マル
	CrossWin                // バツ
	NoWin                   // なんでもない
)
