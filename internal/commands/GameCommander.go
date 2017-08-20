package commands

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"strings"
)

type GameCommander struct {
	AbstractCommander
}

func (commander *GameCommander) ExecuteCommand(commandName string, commandParameters []string) (
	commandResult CommandResult,
) {
	var result string

	command := commander.findCommand(commandName, commandParameters)

	if command == nil {
		result = fmt.Sprintf("Команда \"%v\" не найдена. Список доступен по команде *помощь*.", commandName)
	} else {
		result = command.Execute()
	}

	commandResult = CommandResult{
		Message: tgbotapi.NewMessage(
			commander.connection.ChatId,
			result,
		),
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
