package tests

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/Vehsamrak/telegramud/internal/commands"
)

func TestTest_GetNames_emptyParameters_stringReturned(test *testing.T) {
    command := commands.Test{}

    commandResult := command.Execute()

    assert.Equal(test, commandResult, "Тест пройден!")
}
