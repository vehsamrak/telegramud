package tests

import (
    "testing"
    "github.com/Vehsamrak/telegramud/internal/commands"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/assert"
)

type CommanderFinderMock struct {
    mock.Mock
}

func (mock *CommanderFinderMock) findCommand(requestedCommandName string, commandParameters []string) (
    resultCommand commands.Command,
) {
    return nil
}

func (mock *CommanderFinderMock) createAllCommands(commandParameters []string) []commands.Command {
    return nil
}

func TestMock(test *testing.T) {
    testObj := new(CommanderFinderMock)
    command := commands.Help{Commander: testObj}

    commandResult := command.Execute()

    assert.Equal(
        test,
        commandResult,
        "Доступные команды: ?, chat, help, look, man, test, who, кто, помощь, смотреть, справка, тест, чат",
    )
}
