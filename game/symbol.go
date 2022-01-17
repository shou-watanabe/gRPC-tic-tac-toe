package game

type Symbol int

const (
	Empty  Symbol = iota // 誰も打ってない
	Circle               // マル
	Cross                // バツ
	None                 // なんでもない
)

// 色を文字列に変換します
func SymbolToStr(c Symbol) string {
	switch c {
	case Circle:
		return "○"
	case Cross:
		return "×"
	case Empty:
		return " "
	}

	return ""
}

// 対戦相手の色を取得します
func OpponentSymbol(me Symbol) Symbol {
	switch me {
	case Circle:
		return Cross
	case Cross:
		return Circle
	}

	panic("invalid state")
}
