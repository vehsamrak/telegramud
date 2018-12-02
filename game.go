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

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates, _ := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		var inputText string
		var chatId int64

		if update.Message != nil {
			inputText = update.Message.Text
			chatId = update.Message.Chat.ID
		}

		if update.CallbackQuery != nil {
			inputText = update.CallbackQuery.Data
			chatId = update.CallbackQuery.Message.Chat.ID
			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
		}

		outputMessage := tgbotapi.NewMessage(chatId, inputText)

		switch inputText {
		case "open":
			outputMessage.ReplyMarkup = tgbotapi.NewReplyKeyboard(
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
			)
		case "test":
			outputMessage.Text = "TEST"
			outputMessage.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("test", "ls"),
					tgbotapi.NewInlineKeyboardButtonURL("first link", "google.com"),
					tgbotapi.NewInlineKeyboardButtonURL("second link", "google.com"),
				),
			)
		case "close":
			outputMessage.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		bot.Send(outputMessage)
	}
}
