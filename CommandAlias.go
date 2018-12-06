package main

import "errors"

type CommandAliasParser struct {
}

func (parser *CommandAliasParser) Parse(alias string) (commandName string, err error) {
	aliasMap := map[string]string{
		"Осмотреться":      "look",
		"Заказать выпивку": "drink",
		"Зайти в таверну":  "walk tavern",
		"Выйти на улицу":   "walk street",
		"Портовая улица":   "walk street",
		"Порт":             "walk port",
		"Городские ворота": "walk town_gate",
		"Дорога":           "walk road1",
		"Лесная дорога":    "walk forest_road1",
		"Лесная опушка":    "walk forest1",
		"Полесье":          "walk forest2",
		"Лес":              "walk forest3",
		"Лесная речушка":   "walk forest_river1",
		"Чащоба":           "walk forest4",
	}

	if command, ok := aliasMap[alias]; ok {
		commandName = command
	} else {
		err = errors.New("no alias was found")
	}

	return
}
