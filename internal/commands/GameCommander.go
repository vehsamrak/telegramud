package commands

import (
    "fmt"
    "gopkg.in/telegram-bot-api.v4"
    "strings"

    "github.com/Vehsamrak/telegramud/internal/services"
    "github.com/Vehsamrak/telegramud/internal"
)

type GameCommander struct {
    AbstractCommander
    Messenger *services.Messenger
}

func (commander *GameCommander) ExecuteCommand(commandName string, commandParameters []string) (
    commandResult internal.CommandResult,
) {
    var result string

    command := commander.findCommand(commandName, commandParameters)

    if command == nil {
        result = fmt.Sprintf("Команда \"%v\" не найдена. Список доступен по команде *помощь*.", commandName)
    } else {
        result = command.Execute()
    }

    if result == "" {
        commandResult.IsEmpty = true
    } else {
        commandResult.Message = tgbotapi.NewMessage(commander.connection.ChatId, result)
    }

    return commandResult
}

func (commander *GameCommander) findCommand(requestedCommandName string, commandParameters []string) (resultCommand Command) {
    for _, command := range commander.createAllCommands(commandParameters) {
        for _, commandName := range command.GetNames() {
            if strings.HasPrefix(commandName, requestedCommandName) {
                return command
            }
        }
    }

    return nil
}

// All game commands are created by this method
func (commander *GameCommander) createAllCommands(commandParameters []string) []Command {
    return []Command{
        Test{},
        Look{},
        Who{},
        &Chat{
            Messenger:      commander.Messenger,
            Sender:         *commander.connection,
            ConnectionPool: commander.connectionPool,
            Message:        strings.Join(commandParameters, " "),
        },
        &Help{
            Commander:  commander,
            Parameters: commandParameters,
        },
    }
}