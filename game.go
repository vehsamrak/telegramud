package main

import (
    "log"
    "os"
    "strings"
    "bytes"
    "gopkg.in/telegram-bot-api.v4"

    "github.com/Vehsamrak/telegramud/internal/commands"
)

func main() {
    bot, fault := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APPLICATION_MUD_TOKEN"))
    if fault != nil {
        log.Panic(fault)
    }

    log.Printf("Authorized on account %s", bot.Self.UserName)

    updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 60

    updates, _ := bot.GetUpdatesChan(updateConfig)

    players := map[string]*commands.Connection{}

    for update := range updates {
        if update.Message == nil || update.Message.Text == "" {
            continue
        }

        commandName, commandParameters := parseCommand(update)
        currentUser := players[update.Message.From.UserName]
        executorFactory := commands.ExecutorFactory{ConnectionPool: players}

        var message tgbotapi.MessageConfig
        if currentUser == nil {
            currentUser = &commands.Connection{
                ChatId:   update.Message.Chat.ID,
                UserName: update.Message.From.UserName,
            }

            executor := executorFactory.Create(commands.EXECUTOR_LOGIN, currentUser)
            currentUser.SetExecutor(executor)

            players[update.Message.From.UserName] = currentUser

            message = tgbotapi.NewMessage(
                update.Message.Chat.ID,
                "Добро пожаловать на *Экспериментальный Полигон*!\nИгроки онлайн: " + strings.Join(
                    getPlayersNames(players),
                    ", ",
                ),
            )
        } else {
            commandResult := currentUser.GetExecutor().ExecuteCommand(commandName, commandParameters)
            message = commandResult.Message
            executorName := commandResult.ExecutorName

            if executorName != "" {
                executor := executorFactory.Create(executorName, currentUser)
                currentUser.SetExecutor(executor)
            }
        }

        message.ParseMode = "markdown"
        bot.Send(message)
    }
}

func getPlayersNames(players map[string]*commands.Connection) []string {
    var playerNames []string

    for _, player := range players {
        playerNames = append(playerNames, player.UserName)
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
