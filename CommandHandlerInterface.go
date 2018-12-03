package main

type CommandHandler interface {
	HandleCommand(chatId int64, commandName string) *CommandResult
}
