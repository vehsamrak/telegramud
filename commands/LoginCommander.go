package commands

import (
    "fmt"
    "gopkg.in/telegram-bot-api.v4"
)

type LoginCommander struct {
    Connection *Connection
    Stage string
}

func (commander *LoginCommander) ExecuteCommand(command string, commandParameters []string) (
    commandResult CommandResult,
) {
    switch commander.Stage {
    case "":
        commandResult = CommandResult{
            Message: tgbotapi.NewMessage(
                commander.Connection.ChatId,
                fmt.Sprint("test"),
            ),
        }
    }

    return commandResult
}
