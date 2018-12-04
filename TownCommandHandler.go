package main

type TownCommandHandler struct {
}

func (handler *TownCommandHandler) HandleCommand(commandName string) *CommandResult {
	var command GameCommand

	switch commandName {
	case "tavern":
		command = &TavernCommand{}
	case "drink":
		command = &DrinkCommand{}
	case "street":
		command = &StreetCommand{}
	default:
		command = &UnknownCommand{}
	}

	return command.Execute()
}
