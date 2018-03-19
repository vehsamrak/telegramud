package commands

import (
	"fmt"
	"sort"
	"strings"
)

type Help struct {
	Commander  CommandFinder
	Parameters []string
}

func (command *Help) GetNames() []string {
	return []string{"?", "man", "help", "помощь", "справка"}
}

func (command *Help) Execute() string {
	allCommands := command.Commander.createAllCommands(command.Parameters)

	var allCommandsNames []string
	for _, command := range allCommands {
		allCommandsNames = append(allCommandsNames, command.GetNames()...)
	}

	sort.Strings(allCommandsNames)

	return fmt.Sprint("Доступные команды: ", strings.Join(allCommandsNames, ", "))
}
