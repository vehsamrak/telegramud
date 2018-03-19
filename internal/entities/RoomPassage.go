package entities

type RoomPassage struct {
	RoomFrom *Room
	RoomTo   *Room
	TwoWay   bool
}
