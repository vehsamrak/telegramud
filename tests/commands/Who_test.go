package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vehsamrak/telegramud/internal/commands"
	"testing"
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

	assert.Equal(suite.T(), commandResult, "Игроки онлайн: ")
}
