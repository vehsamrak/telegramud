package tests

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/vehsamrak/telegramud/internal/commands"
    "github.com/stretchr/testify/suite"
    "github.com/vehsamrak/telegramud/internal/entities"
)

func TestLook(test *testing.T) {
    suite.Run(test, new(LookTest))
}

type LookTest struct {
    suite.Suite
}

func (suite *LookTest) SetupTest() {}

func (suite *LookTest) TestGetNames_emptyParameters_stringReturned() {
    command := commands.Look{
        User: &entities.User{
            Room: &entities.Room{
                Name: "Test room name",
                Description: "Test room description",
            },
        },
    }

    commandResult := command.Execute()

    assert.Equal(suite.T(), commandResult, "*Test room name*\nTest room description\n")
}
