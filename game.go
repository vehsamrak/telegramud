package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	playerProvider = PlayerProvider{}.Init()
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APPLICATION_MUD_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	rand.Seed(time.Now().Unix())

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

		if update.CallbackQuery != nil {
			inputText = update.CallbackQuery.Data
			chatId = update.CallbackQuery.Message.Chat.ID
			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
		}

		commandResult := commandHandler.HandleCommand(inputText)

		if commandResult.CommandHandler != nil {
			commandHandler = commandResult.CommandHandler
		}

		output := commandResult.Output()
		output.ChatID = chatId

		bot.Send(output.GenerateChattable())

		for _, command := range commandResult.AfterCommands() {
			commandResult := command.Execute()
			output := commandResult.Output()
			output.ChatID = chatId
			bot.Send(output.GenerateChattable())
		}
	}
}
