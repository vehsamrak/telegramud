package main

import (
    "log"
    "os"
    "strings"
    "gopkg.in/telegram-bot-api.v4"

    "github.com/Vehsamrak/telegramud/models"
    "github.com/Vehsamrak/telegramud/commands"
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

    players := map[string]*models.Connection{}

    for update := range updates {
        if update.Message == nil || update.Message.Text == "" {
            continue
        }

        commandWithParameters := strings.Fields(update.Message.Text)
        commandName := commandWithParameters[0]
        commandParameters := commandWithParameters[1:]

        log.Printf("[%s] %s", update.Message.From.UserName, commandWithParameters)

        currentUser := players[update.Message.From.UserName]
        var message tgbotapi.MessageConfig

        if currentUser == nil {
            currentUser = &models.Connection{
                ChatId: update.Message.Chat.ID,
                UserName: update.Message.From.UserName,
                Stage: "login",
            }

            players[update.Message.From.UserName] = currentUser

            message = tgbotapi.NewMessage(
                update.Message.Chat.ID,
                "Добро пожаловать на Экспериментальный Полигон!\nИгроки онлайн: " + strings.Join(
                    getPlayersNames(players),
                    ", ",
                ),
            )
        } else {
            commander := commands.Commander{Connection: currentUser}
            message = commander.ExecuteCommand(commandName, commandParameters).Message
        }

        bot.Send(message)
    }
}

func getPlayersNames(players map[string]*models.Connection) []string {
    var playerNames []string

    for _, player := range players {
        playerNames = append(playerNames, player.UserName)
    }

    return playerNames
}
