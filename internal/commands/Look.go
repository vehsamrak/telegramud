package commands

import (
    "fmt"

    "github.com/Vehsamrak/telegramud/internal/entities"
)

type Look struct {
    User *entities.User
}

func (command Look) GetNames() []string {
    return []string{"look", "смотреть"}
}

func (command Look) Execute() string {
    currentRoom := command.User.Room

    return fmt.Sprintf("*%v*\n%v\n", currentRoom.Name, currentRoom.Description)
}
