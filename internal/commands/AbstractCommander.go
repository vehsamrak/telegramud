package commands

import "github.com/Vehsamrak/telegramud/internal"

type AbstractCommander struct {
	connection     *internal.Connection
	connectionPool map[string]*internal.Connection
}

func (commander *AbstractCommander) SetConnection(connection *internal.Connection) {
	commander.connection = connection
}

func (commander *AbstractCommander) SetConnectionPool(connectionPool map[string]*internal.Connection) {
	commander.connectionPool = connectionPool
}
