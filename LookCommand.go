package main

import (
	"fmt"
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

	fmt.Printf("%#v\n", "test")

	commandResult.SetOutput(&Output{
		Text: fmt.Sprintf("*%s*\n%s", room.Title(), room.Description()),
	})

	return
}
