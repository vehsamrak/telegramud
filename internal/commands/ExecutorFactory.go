package commands

import (
    "github.com/Vehsamrak/telegramud/internal/services"
    "github.com/Vehsamrak/telegramud/internal"
)

const EXECUTOR_LOGIN string = "login"
const EXECUTOR_GAME string = "game"

type ExecutorFactory struct {
    Messenger      *services.Messenger
    ConnectionPool map[string]*internal.Connection
}

func (factory *ExecutorFactory) Create(executorName string, connection *internal.Connection) (executor internal.Executor) {
    switch executorName {
    case EXECUTOR_GAME:
        executor = &GameCommander{Messenger: factory.Messenger}
    case EXECUTOR_LOGIN:
        executor = &LoginCommander{}
    }

    executor.SetConnection(connection)
    executor.SetConnectionPool(factory.ConnectionPool)

    return executor
}
