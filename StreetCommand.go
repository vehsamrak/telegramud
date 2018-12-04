package main

import "github.com/go-telegram-bot-api/telegram-bot-api"

type StreetCommand struct {
}

func (command *StreetCommand) Name() string {
	return "street"
}

func (command *StreetCommand) Execute() *CommandResult {
	commandResult := &CommandResult{}
	commandResult.SetOutput(&Output{
		Text: "Портовая улица вымощена мраморными плитами, а по обеим сторонам возвышаются колонны. " +
			"На улице многолюдно. Справа виднеется таверна \"*Докер*\".",
		ReplyMarkup: tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Зайти в таверну", "tavern"),
			),
		),
	})

	return commandResult
}
