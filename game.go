package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"bytes"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vehsamrak/telegramud/room"
	"strings"
)

var (
	playerProvider = PlayerProvider{}.Init()
	roomProvider   = room.RoomProvider{}.Init()
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
	commandAliasParser := CommandAliasParser{}.Init()

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

		player := playerProvider.FromTelegramUpdate(update)

		commandName, commandParameters := parseCommand(commandAliasParser, inputText)
		log.Printf("[%s] %s %s", player.Name, commandName, commandParameters)

		commandResult := commandHandler.HandleCommand(player, commandName, commandParameters)

		if commandResult.CommandHandler != nil {
			commandHandler = commandResult.CommandHandler
		}

		output := commandResult.Output(player)
		if output != nil {
			output.ChatID = chatId
			bot.Send(output.GenerateChattable())
		}

		for _, command := range commandResult.AfterCommands() {
			commandResult := command.Execute()
			output := commandResult.Output(player)
			output.ChatID = chatId
			bot.Send(output.GenerateChattable())
		}
	}
}

func parseCommand(
	commandAliasParser *CommandAliasParser,
	rawCommand string,
) (
	commandName string,
	commandParameters []string,
) {
	alias, err := commandAliasParser.Parse(rawCommand)
	if err == nil {
		rawCommand = alias
	}

	rawCommand = strings.TrimSpace(string(bytes.Trim([]byte(rawCommand), "\r\n\x00")))
	commandWithParameters := strings.Fields(rawCommand)

	commandName = strings.ToLower(commandWithParameters[0])

	return commandName, commandWithParameters[1:]
}
