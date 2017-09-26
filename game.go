package main

import (
    "bytes"
    "gopkg.in/telegram-bot-api.v4"
    "log"
    "os"
    "strings"
    "flag"
    "github.com/vehsamrak/osinterrupt"

    "github.com/vehsamrak/telegramud/internal/commands"
    "github.com/vehsamrak/telegramud/internal/services"
    "github.com/vehsamrak/telegramud/internal/entities"
    "github.com/vehsamrak/telegramud/internal/migrations"
)

var (
    database   = &services.Database{}
    worldSaver = &services.WorldSaver{Database: database}
    worldRooms = (&services.RoomGenerator{}).CreateWorld(database.GetConnection())
)

func main() {
    processMigrations()
    worldSaver.Init()

    osinterrupt.HandleTerminateSignal(func() {
        worldSaver.SaveWorld()
        worldSaver.Destruct()
    })

    telegramBot, fault := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APPLICATION_MUD_TOKEN"))
    if fault != nil {
        log.Panic(fault)
    }

    log.Printf("Authorized on account @%s", telegramBot.Self.UserName)

    databaseConnection := database.GetConnection()
    defer worldSaver.SaveWorld()
    defer databaseConnection.Close()

    updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 60

    updates, _ := telegramBot.GetUpdatesChan(updateConfig)

    messenger := &services.Messenger{
        TelegramBot:    telegramBot,
        ConnectionPool: worldSaver.Players,
    }

    for update := range updates {
        if update.Message == nil || update.Message.Text == "" {
            continue
        }

        commandName, commandParameters := parseCommand(update)
        username := update.Message.From.UserName
        currentConnection := worldSaver.Players[username]
        executorFactory := commands.ExecutorFactory{
            ConnectionPool: worldSaver.Players,
            Messenger:      messenger,
            Database:       database,
        }

        var message tgbotapi.MessageConfig
        if currentConnection == nil {
            user := entities.User{UserName: username}
            user = *user.FindByName(database.GetConnection(), username)
            user.Room = worldRooms[0]
            user.ChatId = update.Message.Chat.ID

            currentConnection = &services.Connection{User: &user}

            executor := executorFactory.Create(services.EXECUTOR_LOGIN, currentConnection)
            user.Executor = executor.GetName()
            currentConnection.SetExecutor(executor)

            worldSaver.Players[username] = currentConnection

            messenger.SendMessage(
                update.Message.Chat.ID,
                "Добро пожаловать на *Экспериментальный Полигон*!\nИгроки онлайн: " + strings.Join(
                    getPlayersNames(worldSaver.Players),
                    ", ",
                ),
            )
        } else {
            currentExecutor := currentConnection.GetExecutor()

            if currentExecutor == nil {
                currentExecutor = executorFactory.Create(currentConnection.User.Executor, currentConnection)
            }

            commandResult := currentExecutor.ExecuteCommand(commandName, commandParameters)

            if commandResult.IsEmpty {
                continue
            }

            // if executor was changed by command result
            executorName := commandResult.ExecutorName
            if executorName != "" {
                executor := executorFactory.Create(executorName, currentConnection)
                currentConnection.SetExecutor(executor)
            }

            message = commandResult.Message
            message.ParseMode = "markdown"
            telegramBot.Send(message)
        }
    }
}

// Run automatic migrations on database if flag -migrate passed
func processMigrations() {
    processMigrations := flag.Bool("migrate", false, "Process database migrations")
    flag.Parse()

    if *processMigrations {
        migrations.Migrate()
        os.Exit(0)
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
