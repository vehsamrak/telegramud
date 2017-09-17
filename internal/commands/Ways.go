package commands

import (
    "fmt"
    "strings"

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
        var roomNames []string
        for _, currentRoomPassage := range currentRoomPassages{
            roomNames = append(roomNames, currentRoomPassage.RoomTo.Name)
        }

        commandResult = fmt.Sprintln("Доступные пути:", strings.Join(roomNames, ", "))
    } else {
        commandResult = "Отсюда нет видимых выходов."
    }

    return commandResult
}
