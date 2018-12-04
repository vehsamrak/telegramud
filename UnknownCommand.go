package main

type UnknownCommand struct {
}

func (command *UnknownCommand) Name() string {
	return "unknown"
}

func (command *UnknownCommand) Execute() *CommandResult {
	commandResult := &CommandResult{}
	commandResult.SetOutput(&Output{Text: "Неизвестная команда"})

	return commandResult
}
