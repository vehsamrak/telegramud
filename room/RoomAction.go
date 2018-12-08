package room

type RoomAction struct {
	actionName string
	command    string
}

func (roomExit *RoomAction) ActionName() string {
	return roomExit.actionName
}

func (roomExit *RoomAction) CommandName() string {
	return roomExit.command
}
