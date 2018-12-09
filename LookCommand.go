package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vehsamrak/telegramud/entity"
)

type LookCommand struct {
	commandParameters []string
	player            *entity.Player
}

func (command *LookCommand) Name() string {
	return "look"
}

func (command *LookCommand) SetParameters(commandParameters []string) {
	command.commandParameters = commandParameters
}

func (command *LookCommand) SetPlayer(player *entity.Player) {
	command.player = player
}

func (command *LookCommand) Execute() (commandResult *CommandResult) {
	commandResult = &CommandResult{}

	allRooms := roomProvider.FindAll()
	room := allRooms[command.player.RoomId]

	var keyboard [][]tgbotapi.KeyboardButton
	var actionsKeyboardRow []tgbotapi.KeyboardButton
	var exitsKeyboardRow []tgbotapi.KeyboardButton

	for _, roomAction := range room.Actions() {
		actionButton := tgbotapi.NewKeyboardButton(roomAction.ActionName())
		actionsKeyboardRow = append(actionsKeyboardRow, actionButton)
	}

	if len(actionsKeyboardRow) > 0 {
		keyboard = append(keyboard, actionsKeyboardRow)
	}

	for _, roomExit := range room.Exits() {
		var destinationName string

		if roomExit.DestinationName() == "" {
			destinationName = allRooms[roomExit.DestinationRoomId()].Title()
		} else {
			destinationName = roomExit.DestinationName()
		}

		exitButton := tgbotapi.NewKeyboardButton(destinationName)
		exitsKeyboardRow = append(exitsKeyboardRow, exitButton)
	}

	if len(exitsKeyboardRow) > 0 {
		keyboard = append(keyboard, exitsKeyboardRow)
	}

	commandResult.SetOutput(&Output{
		Text: fmt.Sprintf("*%s*\n%s", room.Title(), room.Description()),
		ReplyMarkup: tgbotapi.ReplyKeyboardMarkup{
			ResizeKeyboard: true,
			Keyboard:       keyboard,
		},
	})

	return
}
