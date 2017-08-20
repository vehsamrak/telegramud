package commands

type Connection struct {
	executor Executor
	ChatId   int64
	UserName string
}

func (connection *Connection) GetExecutor() Executor {
	return connection.executor
}

func (connection *Connection) SetExecutor(executor Executor) {
	connection.executor = executor
}
