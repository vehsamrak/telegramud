package services

import "github.com/vehsamrak/telegramud/internal"

const EXECUTOR_LOGIN string = "login"
const EXECUTOR_GAME string = "game"

type Executor interface {
    ExecuteCommand(command string, commandParameters []string) (commandResult internal.CommandResult)
    SetConnection(connection *Connection)
    SetConnectionPool(connections map[string]*Connection)
    GetName() string
}
