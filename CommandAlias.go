package main

import "errors"

type CommandAliasParser struct {
}

func (parser *CommandAliasParser) Parse(alias string) (commandName string, err error) {
	aliasMap := map[string]string{
		"Осмотреться":      "look",
		"Заказать выпивку": "drink",
	}

	if command, ok := aliasMap[alias]; ok {
		commandName = command
	} else {
		err = errors.New("no alias was found")
	}

	return
}
