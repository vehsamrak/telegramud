package main

import (
	"errors"
)

type CommandAliasParser struct {
}

func (parser *CommandAliasParser) Parse(alias string) (commandName string, err error) {
	aliasMap := map[string]string{
		"Осмотреться": "look",
	}

	allRooms := roomProvider.FindAll()

	// TODO: move to init method
	for _, room := range allRooms {
		aliasMap[room.Title()] = "walk " + room.Id()
		for _, roomAction := range room.Actions() {
			aliasMap[roomAction.ActionName()] = roomAction.CommandName()
		}
		for _, roomExit := range room.Exits() {
			var destinationName string

			if roomExit.DestinationName() == "" {
				destinationName = allRooms[roomExit.DestinationRoomId()].Title()
			} else {
				destinationName = roomExit.DestinationName()
			}

			aliasMap[destinationName] = "walk " + roomExit.DestinationRoomId()
		}
	}

	if command, ok := aliasMap[alias]; ok {
		commandName = command
	} else {
		err = errors.New("no alias was found")
	}

	return
}
