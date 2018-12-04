package main

type DrinkCommand struct {
}

func (command *DrinkCommand) Name() string {
	return "drink"
}

func (command *DrinkCommand) Execute() *CommandResult {
	commandResult := &CommandResult{}
	commandResult.additionalCommands = append(commandResult.additionalCommands, &TavernCommand{})
	commandResult.SetOutput(&Output{
		Text: "Эх, хозяин, пива!",
	})

	return commandResult
}
