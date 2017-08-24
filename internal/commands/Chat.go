package commands

import (
    "fmt"

    "github.com/Vehsamrak/telegramud/internal/services"
    "github.com/Vehsamrak/telegramud/internal"
)

type Chat struct {
    Messenger      *services.Messenger
    Sender         internal.Connection
    ConnectionPool map[string]*internal.Connection
    Message        string
}

func (command *Chat) GetNames() []string {
    return []string{"chat", "talk", "tell", "чат", "говорить"}
}

func (command *Chat) Execute() (result string) {
    chatMessage := fmt.Sprintf("%v говорит: *%v*", command.Sender.User.UserName, command.Message)
    command.Messenger.SendToAll(chatMessage)

    return
}
