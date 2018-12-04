package main

type StartCommandHandler struct {
}

func (handler *StartCommandHandler) HandleCommand(commandName string) *CommandResult {
	commandHandler := &TownCommandHandler{}
	commandResult := &CommandResult{CommandHandler: commandHandler}
	commandResult.additionalCommands = append(commandResult.additionalCommands, &TavernCommand{})

	commandResult.SetOutput(&Output{
		Text: "Добро пожаловать в *Экспериментальный Полигон*!",
	})

	return commandResult
}
