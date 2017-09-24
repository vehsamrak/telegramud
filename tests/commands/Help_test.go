package tests

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "github.com/vehsamrak/telegramud/internal/commands"
    "github.com/vehsamrak/telegramud/internal/services"
)

func TestHelp(test *testing.T) {
    suite.Run(test, new(HelpTest))
}

type HelpTest struct {
    suite.Suite
}

func (suite *HelpTest) SetupTest() {}

func (suite *HelpTest) TestGetNames_emptyParameters_stringReturned() {
    gameCommander := commands.GameCommander{}
    gameCommander.SetConnection(new(services.Connection))
    command := commands.Help{Commander: &gameCommander}

    commandResult := command.Execute()

    assert.Contains(suite.T(), commandResult, "Доступные команды: ?, chat, ")
}
