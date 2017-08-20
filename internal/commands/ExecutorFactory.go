package commands

const EXECUTOR_LOGIN string = "login"
const EXECUTOR_GAME string = "game"

type ExecutorFactory struct {
	ConnectionPool map[string]*Connection
}

func (factory *ExecutorFactory) Create(executorName string, connection *Connection) (executor Executor) {
	switch executorName {
	case EXECUTOR_GAME:
		executor = &GameCommander{}
	case EXECUTOR_LOGIN:
		executor = &LoginCommander{}
	}

	executor.SetConnection(connection)
	executor.SetConnectionPool(factory.ConnectionPool)

	return executor
}
