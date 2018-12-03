package main

type StartCommandHandler struct {
}

func (handler *StartCommandHandler) HandleCommand(commandName string) *CommandResult {
	return &CommandResult{
		Output:         "Добро пожаловать на *Экспериментальный Полигон*!",
		CommandHandler: &TownCommandHandler{},
	}
}
