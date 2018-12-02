package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APPLICATION_MUD_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account @%s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates, _ := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		var inputText string
		var chatId int64
		var callerMessageId int
		var callerMessageText string

		if update.Message != nil {
			inputText = update.Message.Text
			chatId = update.Message.Chat.ID
		}

		if update.CallbackQuery != nil {
			inputText = update.CallbackQuery.Data
			chatId = update.CallbackQuery.Message.Chat.ID
			callerMessageId = update.CallbackQuery.Message.MessageID
			callerMessageText = update.CallbackQuery.Message.Text
			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
		}

		var output *Output
		output = &Output{ChatID: chatId, callerMessageId: callerMessageId, callerMessageText: callerMessageText}

		// input requests
		switch inputText {
		case "open":
			output.SetReplyMarkup(tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("1"),
					tgbotapi.NewKeyboardButton("2"),
					tgbotapi.NewKeyboardButton("3"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("4"),
					tgbotapi.NewKeyboardButton("5"),
					tgbotapi.NewKeyboardButton("6"),
				),
			))
		case "test":
			output.SetText("TEST")
			output.SetReplyMarkup(tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("test", "open"),
					tgbotapi.NewInlineKeyboardButtonURL("first link", "google.com"),
					tgbotapi.NewInlineKeyboardButtonURL("second link", "google.com"),
				),
			))
		case "edit":
			output.SetText("edit pannel")
			output.SetReplyMarkup(tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("edit", "edited"),
				),
			))
		case "close":
			output.SetReplyMarkup(tgbotapi.NewRemoveKeyboard(true))
		}

		// callback requests
		if callerMessageId != 0 {
			switch inputText {
			case "edited":
				if callerMessageId != 0 {
					markup := tgbotapi.NewInlineKeyboardMarkup(
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonData("2", "edited2"),
						),
					)

					output.SetEditKeyboard(&markup)
				}
			case "edited2":
				if callerMessageId != 0 {
					markup := tgbotapi.NewInlineKeyboardMarkup(
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonData("1", "edited"),
						),
					)

					output.SetEditKeyboard(&markup)
				}
			}
		}

		bot.Send(output.CreateChattable())
	}
}
