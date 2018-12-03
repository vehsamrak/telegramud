package main

type TownCommandHandler struct {
}

func (handler *TownCommandHandler) HandleCommand(commandName string) *CommandResult {
	return &CommandResult{
		Output: "Вы находитесь в городе.",
	}
}
