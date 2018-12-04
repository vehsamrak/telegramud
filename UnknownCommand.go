package main

type UnknownCommand struct {
}

func (command *UnknownCommand) Name() string {
	return "unknown"
}

func (command *UnknownCommand) Execute() *CommandResult {
	return &CommandResult{Output: &Output{Text: "Неизвестная команда"}}
}
