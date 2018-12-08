package room

type RoomExit struct {
	destinationRoomId string
	destinationName   string
}

func (roomExit *RoomExit) DestinationRoomId() string {
	return roomExit.destinationRoomId
}

func (roomExit *RoomExit) DestinationName() string {
	return roomExit.destinationName
}
