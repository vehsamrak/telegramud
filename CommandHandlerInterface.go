package main

import "github.com/vehsamrak/telegramud/entity"

type CommandHandler interface {
	HandleCommand(player *entity.Player, commandName string, commandParameters []string) *CommandResult
}
