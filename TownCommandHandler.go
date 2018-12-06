package main

import "github.com/vehsamrak/telegramud/entity"

type TownCommandHandler struct {
}

func (handler *TownCommandHandler) HandleCommand(
	player *entity.Player,
	commandName string,
	commandParameters []string,
) *CommandResult {
	var command GameCommand

	switch commandName {
	case "look":
		command = &LookCommand{}
	case "drink":
		command = &DrinkCommand{}
	case "walk":
		command = &WalkCommand{}
	default:
		command = &UnknownCommand{}
	}

	command.SetParameters(commandParameters)
	command.SetPlayer(player)

	return command.Execute()
}
