package main

import (
	"github.com/vehsamrak/telegramud/entity"
)

type InventoryCommand struct {
	commandParameters []string
	player            *entity.Player
}

func (command *InventoryCommand) Name() string {
	return "inventory"
}

func (command *InventoryCommand) SetParameters(commandParameters []string) {
	command.commandParameters = commandParameters
}

func (command *InventoryCommand) SetPlayer(player *entity.Player) {
	command.player = player
}

func (command *InventoryCommand) Execute() (commandResult *CommandResult) {
	commandResult = &CommandResult{}
	commandResult.SetOutput(&Output{
		Text: "Деньги: 100 $\nПредметы:\nПусто",
	})

	return
}
