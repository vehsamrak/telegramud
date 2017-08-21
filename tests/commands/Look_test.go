package tests

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/Vehsamrak/telegramud/internal/commands"
    "github.com/stretchr/testify/suite"
)

func TestLook(test *testing.T) {
    suite.Run(test, new(LookTest))
}

type LookTest struct {
    suite.Suite
}

func (suite *LookTest) SetupTest() {}

func (suite *LookTest) TestGetNames_emptyParameters_stringReturned() {
    command := commands.Look{}

    commandResult := command.Execute()

    assert.Equal(suite.T(), commandResult, "Ты видишь контуры этого мира.")
}