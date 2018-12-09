package main

type CommandProvider struct {
	commandMap map[string]GameCommand
}

func (provider CommandProvider) Init() *CommandProvider {
	commands := []GameCommand{
		&CharacterCommand{},
		&DrinkCommand{},
		&InventoryCommand{},
		&LookCommand{},
		&WalkCommand{},
	}

	provider.commandMap = map[string]GameCommand{}
	for _, command := range commands {
		provider.commandMap[command.Name()] = command
	}

	return &provider
}

func (provider *CommandProvider) CommandMap() map[string]GameCommand {
	return provider.commandMap
}

func (provider *CommandProvider) FindCommand(commandName string) GameCommand {
	return provider.commandMap[commandName]
}
