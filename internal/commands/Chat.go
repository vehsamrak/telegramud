package commands

import (
    "fmt"

    "github.com/vehsamrak/telegramud/internal/services"
)

type Chat struct {
    Messenger      *services.Messenger
    Sender         services.Connection
    ConnectionPool map[string]*services.Connection
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
