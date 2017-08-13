package commands

type Connection struct {
    ChatId   int64
    UserName string
    Executor Executor
}
