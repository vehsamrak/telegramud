package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vehsamrak/telegramud/entity"
)

type CommandResult struct {
	output             *Output
	CommandHandler     CommandHandler
	additionalCommands []GameCommand
}

func (commandResult *CommandResult) Output(player *entity.Player) *Output {
	output := commandResult.output

	if output != nil && output.ReplyMarkup != nil {
		replyMarkup, ok := output.ReplyMarkup.(tgbotapi.ReplyKeyboardMarkup)

		if ok {
			keyboard := replyMarkup.Keyboard

			characteristicsKeyboardRow := []tgbotapi.KeyboardButton{
				tgbotapi.NewKeyboardButton(
					fmt.Sprintf(
						"Персонаж | %d уровень | %d/%d жизнь",
						player.Level(),
						player.HP(),
						player.MaxHP(),
					),
				),
				tgbotapi.NewKeyboardButton(fmt.Sprintf("Инвентарь | %d $", player.Money())),
			}

			if len(characteristicsKeyboardRow) > 0 {
				keyboard = append(keyboard, characteristicsKeyboardRow)
				replyMarkup.Keyboard = keyboard
				output.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
					ResizeKeyboard: true,
					Keyboard:       keyboard,
				}
			}
		}
	}

	return output
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
