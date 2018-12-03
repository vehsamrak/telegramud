package main

type TownCommandHandler struct {
}

func (handler *TownCommandHandler) HandleCommand(chatId int64, commandName string) *CommandResult {
	return &CommandResult{
		Output: &Output{
			Text:   "Вы находитесь в городе.",
			ChatID: chatId,
		},
	}
}
