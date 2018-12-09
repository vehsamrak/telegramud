package main

import (
	"fmt"
	"github.com/vehsamrak/telegramud/entity"
)

type CharacterCommand struct {
	commandParameters []string
	player            *entity.Player
}

func (command *CharacterCommand) Name() string {
	return "character"
}

func (command *CharacterCommand) SetParameters(commandParameters []string) {
	command.commandParameters = commandParameters
}

func (command *CharacterCommand) SetPlayer(player *entity.Player) {
	command.player = player
}

func (command *CharacterCommand) Execute() (commandResult *CommandResult) {
	commandResult = &CommandResult{}
	commandResult.SetOutput(&Output{
		Text: fmt.Sprintf(
			"Имя: %s\nУровень: 1\n"+
				"Раса: человек\n"+
				"Профессия: не выбрана\n"+
				"Здоровье: 100 / 100",
			command.player.Name,
		),
	})

	return
}
