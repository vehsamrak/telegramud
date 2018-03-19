package commands

import (
	"strings"

	"github.com/vehsamrak/telegramud/internal/entities"
)

type Go struct {
	User              *entities.User
	CommandParameters []string
}

func (command Go) GetNames() []string {
	return []string{"go", "идти", "перейти", "пройти", "пойти"}
}

func (command Go) Execute() (commandResult string) {
	currentRoomPassages := command.User.Room.Passages

	if currentRoomPassages != nil {
		var destinationRoomName string

		if len(command.CommandParameters) > 0 {
			destinationRoomName = command.CommandParameters[0]
		} else {
			commandResult = "Куда ты хочешь пойти?"
			return
		}

		var roomNames []string
		roomReached := false
		for _, currentRoomPassage := range currentRoomPassages {
			roomNames = append(roomNames, currentRoomPassage.RoomTo.Name)
			if strings.HasPrefix(strings.ToLower(currentRoomPassage.RoomTo.Name), strings.ToLower(destinationRoomName)) {
				command.User.Room = currentRoomPassage.RoomTo
				roomReached = true
				break
			}
		}

		if roomReached {
			commandResult = "Ты дошел до места назначения.\n\n" + Look{User: command.User}.Execute()
		} else {
			commandResult = "Ты не можешь пойти туда."
		}
	} else {
		commandResult = Ways{User: command.User}.Execute()
	}

	return commandResult
}
