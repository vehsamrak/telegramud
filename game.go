package main

import (
    "log"
    "os"

    "gopkg.in/telegram-bot-api.v4"
    "strings"
    "fmt"
)

func main() {
    bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APPLICATION_MUD_TOKEN"))
    if err != nil {
        log.Panic(err)
    }

    log.Printf("Authorized on account %s", bot.Self.UserName)

    update := tgbotapi.NewUpdate(0)
    update.Timeout = 60

    updates, err := bot.GetUpdatesChan(update)

    players := map[string]*Connection{}

    for update := range updates {
        if update.Message == nil {
            continue
        }

        commandWithParameters := strings.Fields(update.Message.Text)
        commandName := commandWithParameters[0]
        commandParameters := commandWithParameters[1:]

        log.Printf("[%s] %s", update.Message.From.UserName, commandWithParameters)

        currentUser := players[update.Message.From.UserName]

        if currentUser == nil {
            currentUser = &Connection{
                ChatId: update.Message.Chat.ID,
                UserName: update.Message.From.UserName,
                Stage: "login",
            }

            players[update.Message.From.UserName] = currentUser

            msg := tgbotapi.NewMessage(
                update.Message.Chat.ID,
                "Привет!\nИгроки онлайн: " + strings.Join(getPlayersNames(players), ", "),
            )

            bot.Send(msg)
        } else {
            msg := tgbotapi.NewMessage(
                update.Message.Chat.ID,
                fmt.Sprintf("Введена команда: \"%v\". Параметры: %v", commandName, commandParameters),
            )

            bot.Send(msg)
        }
    }
}

func getPlayersNames(players map[string]*Connection) []string {
    var playerNames []string

    for _, player := range players {
        playerNames = append(playerNames, player.UserName)
    }

    return playerNames
}
