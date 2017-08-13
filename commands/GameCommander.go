package commands

import (
    "fmt"
    "gopkg.in/telegram-bot-api.v4"
)

type GameCommander struct {
    AbstractCommander
}

func (commander *GameCommander) ExecuteCommand(command string, commandParameters []string) (
    commandResult CommandResult,
) {
    var result string

    switch command {
    case "test":
        result = "Passed"
    default:
        result = fmt.Sprintf("Команда \"%v\" не найдена.", command)
    }

    commandResult = CommandResult{
        Message: tgbotapi.NewMessage(
            commander.connection.ChatId,
            result,
        ),
    }

    return commandResult
}
