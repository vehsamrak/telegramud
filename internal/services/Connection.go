package services

import "github.com/Vehsamrak/telegramud/internal/entities"

type Connection struct {
    executor Executor
    ChatId   int64
    User     entities.User
}

func (connection *Connection) GetExecutor() Executor {
    return connection.executor
}

func (connection *Connection) SetExecutor(executor Executor) {
    connection.executor = executor
}
