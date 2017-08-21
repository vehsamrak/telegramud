package commands

type CommandFinder interface {
    findCommand(requestedCommandName string, commandParameters []string) (resultCommand Command)
    createAllCommands(commandParameters []string) []Command
}
