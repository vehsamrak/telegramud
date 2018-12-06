package main

import "github.com/vehsamrak/telegramud/entity"

type WalkCommand struct {
	commandParameters []string
	player            *entity.Player
}

func (command *WalkCommand) Name() string {
	return "walk"
}

func (command *WalkCommand) SetParameters(commandParameters []string) {
	command.commandParameters = commandParameters
}

func (command *WalkCommand) SetPlayer(player *entity.Player) {
	command.player = player
}

func (command *WalkCommand) Execute() (commandResult *CommandResult) {
	commandResult = &CommandResult{}

	if len(command.commandParameters) == 0 {
		unknownCommand := &UnknownCommand{}
		commandResult = unknownCommand.Execute()
		return
	}

	destination := command.commandParameters[0]
	room := roomProvider.FindById(destination)

	if room == nil {
		commandResult.SetOutput(&Output{
			Text: "Туда не пройти.",
		})
		return
	}

	command.player.RoomId = room.Id()

	lookCommand := &LookCommand{player: command.player}
	commandResult.AddCommand(lookCommand)

	return commandResult
}
