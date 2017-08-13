package commands

import (
    "fmt"
    "gopkg.in/telegram-bot-api.v4"

    "github.com/Vehsamrak/telegramud/models"
)

type Commander struct {
	Connection *models.Connection
}

func (commander *Commander) ExecuteCommand(command string, commandParameters []string) (commandResult CommandResult) {
	if command == "" {
		return
	}

    if commander.Connection.Stage == "login" {
        return CommandResult{
            Message: tgbotapi.NewMessage(
                commander.Connection.ChatId,
                fmt.Sprintf("Введена команда: \"%v\". Параметры: %v", command, commandParameters),
            ),
        }
    } else {
        return
    }

	return commandResult
}
