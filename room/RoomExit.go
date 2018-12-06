package room

type RoomAction struct {
	commandName string
}

func (roomExit *RoomAction) CommandName() string {
	return roomExit.commandName
}
