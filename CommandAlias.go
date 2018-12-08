package main

import (
	"errors"
)

type CommandAliasParser struct {
	aliasMap map[string]string
}

func (parser CommandAliasParser) Init() *CommandAliasParser {
	aliasMap := map[string]string{
		"Осмотреться": "look",
	}

	allRooms := roomProvider.FindAll()

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

	parser.aliasMap = aliasMap

	return &parser
}

func (parser *CommandAliasParser) Parse(alias string) (commandName string, err error) {
	if command, ok := parser.aliasMap[alias]; ok {
		commandName = command
	} else {
		err = errors.New("no alias was found")
	}

	return
}
