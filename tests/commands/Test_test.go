package tests

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/Vehsamrak/telegramud/internal/commands"
    "github.com/stretchr/testify/suite"
)

func Test_Test(test *testing.T) {
    suite.Run(test, new(TestTest))
}

type TestTest struct {
    suite.Suite
}

func (suite *TestTest) SetupTest() {}

func (suite *TestTest) TestGetNames_emptyParameters_stringReturned() {
    command := commands.Test{}

    commandResult := command.Execute()

    assert.Equal(suite.T(), commandResult, "Тест пройден!")
}