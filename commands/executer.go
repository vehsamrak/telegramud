package commands

type Executor interface {
	ExecuteCommand(command string) (commandResult CommandResult)
}
