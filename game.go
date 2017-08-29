package main

import (
    "bytes"
    "gopkg.in/telegram-bot-api.v4"
    "log"
    "os"
    "strings"

    "github.com/Vehsamrak/telegramud/internal/commands"
    "github.com/Vehsamrak/telegramud/internal/services"
    "github.com/Vehsamrak/telegramud/internal/entities"
)

var (
    database = &services.Database{}
    worldRooms = (&services.RoomGenerator{}).GenerateWorld()
)

func main() {
    telegramBot, fault := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APPLICATION_MUD_TOKEN"))
    if fault != nil {
        log.Panic(fault)
    }

    log.Printf("Authorized on account @%s", telegramBot.Self.UserName)

    databaseConnection := database.GetConnection()
    defer databaseConnection.Close()

    updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 60

    updates, _ := telegramBot.GetUpdatesChan(updateConfig)

    players := map[string]*services.Connection{}
    messenger := &services.Messenger{
        TelegramBot:    telegramBot,
        ConnectionPool: players,
    }

    for update := range updates {
        if update.Message == nil || update.Message.Text == "" {
            continue
        }

        commandName, commandParameters := parseCommand(update)
        username := update.Message.From.UserName
        currentConnection := players[username]
        executorFactory := commands.ExecutorFactory{
            ConnectionPool: players,
            Messenger:      messenger,
            Database:       database,
        }

        var message tgbotapi.MessageConfig
        if currentConnection == nil {
            user := entities.User{UserName: username}
            user = *user.FindByName(database.GetConnection(), username)
            user.Room = worldRooms[0]

            currentConnection = &services.Connection{
                ChatId: update.Message.Chat.ID,
                User:   &user,
            }

            executor := executorFactory.Create(commands.EXECUTOR_LOGIN, currentConnection)
            currentConnection.SetExecutor(executor)

            players[username] = currentConnection

            messenger.SendMessage(
                update.Message.Chat.ID,
                "Добро пожаловать на *Экспериментальный Полигон*!\nИгроки онлайн: " + strings.Join(
                    getPlayersNames(players),
                    ", ",
                ),
            )
        } else {
            commandResult := currentConnection.GetExecutor().ExecuteCommand(commandName, commandParameters)
            message = commandResult.Message
            executorName := commandResult.ExecutorName

            if commandResult.IsEmpty {
                continue
            }

            if executorName != "" {
                executor := executorFactory.Create(executorName, currentConnection)
                currentConnection.SetExecutor(executor)
            }

            message.ParseMode = "markdown"
            telegramBot.Send(message)
        }
    }
}

func getPlayersNames(players map[string]*services.Connection) []string {
    var playerNames []string

    for _, player := range players {
        playerNames = append(playerNames, player.User.UserName)
    }

    return playerNames
}

func parseCommand(update tgbotapi.Update) (commandName string, commandParameters []string) {
    rawCommand := string(bytes.Trim([]byte(update.Message.Text), "\r\n\x00"))
    rawCommand = strings.TrimSpace(rawCommand)
    commandWithParameters := strings.Fields(rawCommand)

    log.Printf("[%s] %s", update.Message.From.UserName, commandWithParameters)

    commandName = strings.ToLower(commandWithParameters[0])

    return commandName, commandWithParameters[1:]
}
