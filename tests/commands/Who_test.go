package tests

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "github.com/Vehsamrak/telegramud/internal/commands"
)

func TestWho(test *testing.T) {
    suite.Run(test, new(WhoTest))
}

type WhoTest struct {
    suite.Suite
}

func (suite *WhoTest) SetupTest() {}

func (suite *WhoTest) TestGetNames_emptyParameters_stringReturned() {
    command := commands.Who{}

    commandResult := command.Execute()

    assert.Equal(suite.T(), commandResult, "В этом мире нет никого лучше тебя.")
}
