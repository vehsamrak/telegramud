package entity

type Player struct {
	name         string
	ChatId       int64
	RoomId       string
	healthPoints int32
	money        int32
	level        int32
}

func (player Player) Init() *Player {
	player.money = 10
	player.level = 1
	player.healthPoints = 10
	return &player
}

func (player *Player) Name() string {
	return player.name
}

func (player *Player) HP() int32 {
	return player.healthPoints
}

func (player *Player) MaxHP() int32 {
	return player.level * 10
}

func (player *Player) Money() int32 {
	return player.money
}

func (player *Player) Level() int32 {
	return player.level
}
func (player *Player) SetName(username string) {
	player.name = username
}
