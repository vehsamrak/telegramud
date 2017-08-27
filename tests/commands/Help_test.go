package tests

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "github.com/Vehsamrak/telegramud/internal/commands"
    "github.com/Vehsamrak/telegramud/internal/services"
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

    assert.Equal(suite.T(), commandResult, "2Доступные команды: ?, chat, help, look, man, talk, tell, test, who, говорить, кто, помощь, смотреть, справка, тест, чат")
}
