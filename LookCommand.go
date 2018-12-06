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

	room := roomProvider.FindByPlayer(command.player)

	var exitKeyboard [][]tgbotapi.KeyboardButton
	var exitKeyboardRow []tgbotapi.KeyboardButton

	for _, roomAction := range room.Actions() {
		exitButton := tgbotapi.NewKeyboardButton(roomAction.CommandTitle())
		exitKeyboardRow = append(exitKeyboardRow, exitButton)
	}

	exitKeyboard = append(exitKeyboard, exitKeyboardRow)

	commandResult.SetOutput(&Output{
		Text: fmt.Sprintf("*%s*\n%s", room.Title(), room.Description()),
		ReplyMarkup: tgbotapi.ReplyKeyboardMarkup{
			ResizeKeyboard: true,
			Keyboard:       exitKeyboard,
		},
	})

	return
}
