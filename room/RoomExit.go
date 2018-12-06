package room

type RoomAction struct {
	commandName  string
	commandTitle string
}

func (roomExit *RoomAction) CommandName() string {
	return roomExit.commandName
}

func (roomExit *RoomAction) CommandTitle() string {
	return roomExit.commandTitle
}
