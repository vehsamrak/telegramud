package internal

import "gopkg.in/telegram-bot-api.v4"

type CommandResult struct {
    IsEmpty      bool
    Message      tgbotapi.MessageConfig
    ExecutorName string
}
