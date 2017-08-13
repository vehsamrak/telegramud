package commands

import (
    "fmt"
    "gopkg.in/telegram-bot-api.v4"
)

type LoginCommander struct {
    Connection *Connection
    Stage      string
    Choice     ChoiceControl
}

func (commander *LoginCommander) ExecuteCommand(command string, commandParameters []string) (
    commandResult CommandResult,
) {
    var result string

    if commander.Choice.Question != "" {
        commander.Choice.Answer = command

        if commander.Choice.CheckAnswer() {
            commander.Stage = commander.Choice.AfterStage
            commander.Choice = ChoiceControl{}
        } else {
            result = commander.Choice.GetQuestionMessage()
        }
    }

    if result == "" {
        switch commander.Stage {
        case "":
            commander.Choice = ChoiceControl{
                Question: "Сделайте свой выбор",
                AvailableAnswers: []string{"один", "два", "три"},
                AfterStage: "2",
            }

            result = commander.Choice.GetQuestionMessage()
        case "2":
            result = fmt.Sprint("Выбор сделан.")
        }
    }

    commandResult = CommandResult{
        Message: tgbotapi.NewMessage(
            commander.Connection.ChatId,
            result,
        ),
    }

    return commandResult
}
