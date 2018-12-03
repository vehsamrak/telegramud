package main

import "github.com/go-telegram-bot-api/telegram-bot-api"

type TownCommandHandler struct {
}

func (handler *TownCommandHandler) HandleCommand(chatId int64, commandName string) *CommandResult {
	return &CommandResult{
		Output: &Output{
			Text: "Вы находитесь в шумной городской таверне.",
			ReplyMarkup: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Заказать выпивку", "drink"),
					tgbotapi.NewInlineKeyboardButtonData("Выйти на улицу", "exit"),
				),
			),
			ChatID: chatId,
		},
	}
}
