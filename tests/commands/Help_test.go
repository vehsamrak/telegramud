package tests

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "github.com/Vehsamrak/telegramud/internal/commands"
)

func TestHelp(test *testing.T) {
    suite.Run(test, new(HelpTest))
}

type HelpTest struct {
    suite.Suite
    VariableThatShouldStartAtFive int
}

func (suite *HelpTest) SetupTest() {}

func (suite *HelpTest) TestGetNames_emptyParameters_stringReturned() {
    gameCommander := commands.GameCommander{}
    gameCommander.SetConnection(new(commands.Connection))
    command := commands.Help{Commander: &gameCommander}

    commandResult := command.Execute()

    assert.Equal(suite.T(), commandResult, "Доступные команды: ?, chat, help, look, man, test, who, кто, помощь, смотреть, справка, тест, чат")
}
