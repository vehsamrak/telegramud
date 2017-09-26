package services

import (
    "gopkg.in/telegram-bot-api.v4"
)

type Messenger struct {
    TelegramBot    *tgbotapi.BotAPI
    ConnectionPool map[string]*Connection
}

func (sender *Messenger) SendMessage(chatId int64, message string) {
    telegramMessage := tgbotapi.NewMessage(chatId, message)
    telegramMessage.ParseMode = "markdown"

    sender.TelegramBot.Send(telegramMessage)
}

func (sender *Messenger) SendToAll(message string) {
    for _, connection := range sender.ConnectionPool {
        sender.SendMessage(connection.User.ChatId, message)
    }
}
