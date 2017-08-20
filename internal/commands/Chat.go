package commands

import (
	"fmt"
)

type Chat struct {
	Sender         Connection
	ConnectionPool map[string]*Connection
	Message        string
}

func (command *Chat) GetNames() []string {
	return []string{"chat", "чат"}
}

func (command *Chat) Execute() string {
	return fmt.Sprintf("%v говорит: *%v*", command.Sender.UserName, command.Message)
}
