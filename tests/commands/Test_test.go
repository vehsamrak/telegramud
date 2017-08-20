package tests

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/Vehsamrak/telegramud/commands"
)

func TestGetNames(test *testing.T) {
    testCommand := commands.Test{}

    commandResult := testCommand.Execute()

    assert.Equal(test, commandResult, "Тест пройден!")
}
