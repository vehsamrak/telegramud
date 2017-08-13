package commands

import "gopkg.in/telegram-bot-api.v4"

type CommandResult struct {
    Message      tgbotapi.MessageConfig
    ExecutorName string
}
