package main

import (
	"github.com/vehsamrak/telegramud/entity"
)

type StartCommandHandler struct {
}

func (handler *StartCommandHandler) HandleCommand(
	player *entity.Player,
	commandName string,
	commandParameters []string,
) *CommandResult {
	commandResult := &CommandResult{CommandHandler: &TownCommandHandler{}}

	commandResult.SetOutput(&Output{
		Text: "Добро пожаловать в *Портовый Город*!",
	})
	commandResult.AddCommand(&LookCommand{player: player})

	return commandResult
}
