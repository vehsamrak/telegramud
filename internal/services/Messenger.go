package services

import (
    "gopkg.in/telegram-bot-api.v4"

    "github.com/Vehsamrak/telegramud/internal"
)

type Messenger struct {
    TelegramBot    *tgbotapi.BotAPI
    ConnectionPool map[string]*internal.Connection
}

func (sender *Messenger) SendMessage(chatId int64, message string) {
    telegramMessage := tgbotapi.NewMessage(chatId, message)
    telegramMessage.ParseMode = "markdown"

    sender.TelegramBot.Send(telegramMessage)
}

func (sender *Messenger) SendToAll(message string) {
    for _, connection := range sender.ConnectionPool {
        sender.SendMessage(connection.ChatId, message)
    }
}
