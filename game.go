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

	introUpdateConfig := tgbotapi.NewUpdate(0)
	introUpdateConfig.Timeout = 60
	introUpdates, _ := bot.GetUpdatesChan(introUpdateConfig)

	var commandHandler CommandHandler
	commandHandler = &StartCommandHandler{}

	for update := range introUpdates {
		if update.Message == nil && update.CallbackQuery == nil {
			continue
		}

		var inputText string
		var chatId int64

		if update.Message != nil {
			inputText = update.Message.Text
			chatId = update.Message.Chat.ID
		}

		commandResult := commandHandler.HandleCommand(chatId, inputText)

		if commandResult.CommandHandler != nil {
			commandHandler = commandResult.CommandHandler
		}

		output := commandResult.Output

		bot.Send(output.CreateChattable())
	}
}
