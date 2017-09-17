package commands

import (
    "github.com/Vehsamrak/telegramud/internal/entities"
)

type Go struct {
    User *entities.User
}

func (command Go) GetNames() []string {
	return []string{"go", "идти", "перейти", "пройти", "пойти"}
}

func (command Go) Execute() string {
    currentRoomPassages := command.User.Room.Passages

    var commandResult string

    if currentRoomPassages != nil {
        commandResult = Look{command.User}.Execute()
    } else {
        commandResult = "Ты не можешь пройти в том направлении."
    }

    return commandResult
}
