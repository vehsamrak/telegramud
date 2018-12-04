package main

import "github.com/go-telegram-bot-api/telegram-bot-api"

type LookCommand struct {
}

func (command *LookCommand) Name() string {
	return "look"
}

func (command *LookCommand) Execute() *CommandResult {
	return &CommandResult{Output: &Output{
		Text: "Вы находитесь в шумной городской таверне.",
		ReplyMarkup: tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Заказать выпивку", "drink"),
				tgbotapi.NewInlineKeyboardButtonData("Выйти на улицу", "exit"),
			),
		),
	}}
}
