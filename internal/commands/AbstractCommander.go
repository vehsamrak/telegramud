package commands

import (
	"github.com/vehsamrak/telegramud/internal/services"
)

type AbstractCommander struct {
	connection     *services.Connection
	connectionPool map[string]*services.Connection
}

func (commander *AbstractCommander) SetConnection(connection *services.Connection) {
	commander.connection = connection
}

func (commander *AbstractCommander) SetConnectionPool(connectionPool map[string]*services.Connection) {
	commander.connectionPool = connectionPool
}
