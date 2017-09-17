package commands

import (
    "github.com/Vehsamrak/telegramud/internal/entities"
)

type Ways struct {
    User *entities.User
}

func (command Ways) GetNames() []string {
	return []string{"ways", "exits", "пути", "выходы"}
}

func (command Ways) Execute() string {
    currentRoomPassages := command.User.Room.Passages

    var commandResult string

    if currentRoomPassages != nil {
    } else {
        commandResult = "Отсюда нет видимых выходов."
    }

    return commandResult
}
