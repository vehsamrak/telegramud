package services

import "github.com/Vehsamrak/telegramud/internal"

type Executor interface {
    ExecuteCommand(command string, commandParameters []string) (commandResult internal.CommandResult)
    SetConnection(connection *Connection)
    SetConnectionPool(connections map[string]*Connection)
}