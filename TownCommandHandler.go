package main

type TownCommandHandler struct {
}

func (handler *TownCommandHandler) HandleCommand(commandName string) *CommandResult {
	var command GameCommand

	switch commandName {
	case "look":
		command = &LookCommand{}
	default:
		command = &UnknownCommand{}
	}

	return command.Execute()
}
