package commands

import "github.com/vehsamrak/telegramud/internal/services"

type ExecutorFactory struct {
    Messenger      *services.Messenger
    ConnectionPool map[string]*services.Connection
    Database       *services.Database
}

func (factory *ExecutorFactory) Create(executorName string, connection *services.Connection) (executor services.Executor) {
    switch executorName {
    case services.EXECUTOR_GAME:
        executor = &GameCommander{
            Messenger: factory.Messenger,
            Database:  factory.Database,
        }
    case services.EXECUTOR_LOGIN:
        executor = &LoginCommander{Database: factory.Database}
    }

    executor.SetConnection(connection)
    executor.SetConnectionPool(factory.ConnectionPool)

    return executor
}
