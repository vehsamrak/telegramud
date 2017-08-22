package commands

import (
    "strings"

    "github.com/Vehsamrak/telegramud/internal"
)

type Who struct {
    ConnectionPool map[string]*internal.Connection
}

func (command *Who) GetNames() []string {
    return []string{"who", "кто"}
}

func (command *Who) Execute() string {
    return "Игроки онлайн: " + strings.Join(command.getOnlinePlayersNames(), ", ")
}

func (command *Who) getOnlinePlayersNames() (playerNames []string) {
    for _, player := range command.ConnectionPool {
        playerNames = append(playerNames, player.UserName)
    }

    return
}
