package main

import (
	"fmt"
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
		Text: fmt.Sprintf("Деньги: %d $\nПредметы:\nПусто", command.player.Money()),
	})

	return
}
