package main

import "github.com/vehsamrak/telegramud/entity"

type UnknownCommand struct {
	commandParameters []string
	player            *entity.Player
}

func (command *UnknownCommand) Name() string {
	return "unknown"
}

func (command *UnknownCommand) SetParameters(commandParameters []string) {
	command.commandParameters = commandParameters
}

func (command *UnknownCommand) SetPlayer(player *entity.Player) {
	command.player = player
}

func (command *UnknownCommand) Execute() *CommandResult {
	commandResult := &CommandResult{}
	commandResult.SetOutput(&Output{Text: "Неизвестная команда"})

	return commandResult
}
