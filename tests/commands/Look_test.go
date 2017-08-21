package tests

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/Vehsamrak/telegramud/internal/commands"
)

func TestLook_GetNames_emptyParameters_stringReturned(test *testing.T) {
    command := commands.Look{}

    commandResult := command.Execute()

    assert.Equal(test, commandResult, "Ты видишь контуры этого мира.")
}
