package main

type CommandResult struct {
	output             *Output
	CommandHandler     CommandHandler
	additionalCommands []GameCommand
}

func (commandResult *CommandResult) Output() *Output {
	return commandResult.output
}

func (commandResult *CommandResult) SetOutput(output *Output) {
	commandResult.output = output
}

func (commandResult *CommandResult) AfterCommands() []GameCommand {
	return commandResult.additionalCommands
}

func (commandResult *CommandResult) AddCommand(command GameCommand) {
	commandResult.additionalCommands = append(commandResult.additionalCommands, command)
}
