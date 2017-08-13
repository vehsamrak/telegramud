package commands

type Executor interface {
    ExecuteCommand(command string, commandParameters []string) (commandResult CommandResult)
    SetConnection(connection *Connection)
}
