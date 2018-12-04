package main

type StartCommandHandler struct {
}

func (handler *StartCommandHandler) HandleCommand(commandName string) *CommandResult {
	commandHandler := &TownCommandHandler{}
	commandResult := &CommandResult{CommandHandler: commandHandler}
	commandResult.additionalCommands = append(commandResult.additionalCommands, &LookCommand{})

	commandResult.SetOutput(&Output{
		Text: "Добро пожаловать на *Экспериментальный Полигон*!",
	})

	return commandResult
}
