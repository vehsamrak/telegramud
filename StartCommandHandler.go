package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
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
		Text: "Добро пожаловать в *Экспериментальный Полигон*!",
		ReplyMarkup: tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Осмотреться"),
				tgbotapi.NewKeyboardButton("Заказать выпивку"),
			),
		),
	})

	return commandResult
}
