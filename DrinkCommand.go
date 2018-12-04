package main

import "math/rand"

type DrinkCommand struct {
}

func (command *DrinkCommand) Name() string {
	return "drink"
}

func (command *DrinkCommand) Execute() *CommandResult {
	commandResult := &CommandResult{}
	commandResult.additionalCommands = append(commandResult.additionalCommands, &TavernCommand{})
	commandResult.SetOutput(&Output{
		Text: command.generateTavernPhrase(),
	})

	return commandResult
}

func (command *DrinkCommand) generateTavernPhrase() string {
	mobNames := []string{
		"Вы выпили освежающего пива.",
		"Вы промочили горло ромом.",
		"Вы выпили непонятную мутную алкогольную жидкость.",
		"Допив стакан джина Вы осознали что пьяны.",
		"Вы попытались выпить еще, но упали со стула.",
		"Выпив, Вы подумали: \"_пить - здоровье лечить_\".",
		"Выпив, Вы подумали: \"_Все у меня хорошо. А если не совсем, то без проблем - заказываю пиво еще!_\".",
		"Вы выпили, а затем выпили еще.",
		"Вы освежились пенным напитком.",
		"Залпом опрокинув стакан рома Вы твердо решили выпить еще.",
	}

	n := rand.Int() % len(mobNames)

	return mobNames[n]
}
