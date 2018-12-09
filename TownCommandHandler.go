package main

import (
	"github.com/vehsamrak/telegramud/entity"
)

type CommandHandler struct {
}

func (handler *CommandHandler) HandleCommand(
	player *entity.Player,
	commandName string,
	commandParameters []string,
) *CommandResult {
	commandProvider := CommandProvider{}.Init()
	command := commandProvider.FindCommand(commandName)

	if command == nil {
		command = &UnknownCommand{}
	}

	command.SetParameters(commandParameters)
	command.SetPlayer(player)

	return command.Execute()
}
