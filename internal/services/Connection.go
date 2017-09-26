package services

import "github.com/vehsamrak/telegramud/internal/entities"

type Connection struct {
    executor Executor
    User     *entities.User
}

func (connection *Connection) GetExecutor() Executor {
    return connection.executor
}

func (connection *Connection) SetExecutor(executor Executor) {
    connection.executor = executor
    connection.User.Executor = executor.GetName()
}
