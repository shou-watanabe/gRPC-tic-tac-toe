package entity

type Room struct {
	ID    int32
	Host  *Player
	Guest *Player
}
