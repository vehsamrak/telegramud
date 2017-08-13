package commands

type AbstractCommander struct {
    connection *Connection
    connectionPool map[string]*Connection
}

func (commander *AbstractCommander) SetConnection(connection *Connection) {
    commander.connection = connection
}

func (commander *AbstractCommander) SetConnectionPool(connectionPool map[string]*Connection) {
    commander.connectionPool = connectionPool
}
