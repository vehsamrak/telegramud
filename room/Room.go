package room

type Room struct {
	id          string
	title       string
	description string
	actions     []*RoomAction
}

func (room *Room) Id() string {
	return room.id
}

func (room *Room) Title() string {
	return room.title
}

func (room *Room) Description() string {
	return room.description
}

func (room *Room) Actions() []*RoomAction {
	return room.actions
}
