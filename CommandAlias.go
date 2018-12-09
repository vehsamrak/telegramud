package main

import (
	"errors"
	"strings"
)

type CommandAliasParser struct {
	aliasMap map[string]string
}

func (parser CommandAliasParser) Init() *CommandAliasParser {
	aliasMap := map[string]string{}

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

func (parser *CommandAliasParser) Parse(input string) (commandName string, err error) {
	if command, ok := parser.aliasMap[input]; ok {
		commandName = command
	} else if true {
		aliases := map[string]string{
			"Осмотреться": "look",
			"Персонаж":    "character",
			"Инвентарь":   "inventory",
		}

		for aliasName, aliasReferenceName := range aliases {
			inputCommandWithParameters := strings.Fields(input)
			inputCommand := inputCommandWithParameters[0]

			if inputCommand == aliasName {
				commandName = aliasReferenceName
			}
		}
	} else {
		err = errors.New("no alias was found")
	}

	return
}
