package main

type CommandHandler interface {
	HandleCommand(commandName string) *CommandResult
}
