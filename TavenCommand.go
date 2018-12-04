package main

import "github.com/go-telegram-bot-api/telegram-bot-api"

type TavernCommand struct {
}

func (command *TavernCommand) Name() string {
	return "tavern"
}

func (command *TavernCommand) Execute() *CommandResult {
	commandResult := &CommandResult{}
	commandResult.SetOutput(&Output{
		Text: "Вы находитесь в шумной городской таверне.",
		ReplyMarkup: tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Заказать выпивку", "drink"),
				tgbotapi.NewInlineKeyboardButtonData("Выйти на улицу", "street"),
			),
		),
	})

	return commandResult
}
