package entity

type Symbol int

const (
	Empty  Symbol = iota // 誰も打ってない
	Circle               // マル
	Cross                // バツ
	None                 // なんでもない
)
