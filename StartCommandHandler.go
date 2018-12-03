package main

type StartCommandHandler struct {
}

func (handler *StartCommandHandler) HandleCommand(chatId int64, commandName string) *CommandResult {
	return &CommandResult{
		Output: &Output{
			Text:   "Добро пожаловать на *Экспериментальный Полигон*!",
			ChatID: chatId,
		},
		CommandHandler: &TownCommandHandler{},
	}
}
