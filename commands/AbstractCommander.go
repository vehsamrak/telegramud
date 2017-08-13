package commands

type AbstractCommander struct {
    connection *Connection
}

func (commander *AbstractCommander) SetConnection(connection *Connection) {
    commander.connection = connection
}
