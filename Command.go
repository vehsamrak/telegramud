package main

import "github.com/vehsamrak/telegramud/entity"

type GameCommand interface {
	Name() string
	Execute() *CommandResult
	SetParameters(commandParameters []string)
	SetPlayer(player *entity.Player)
}
