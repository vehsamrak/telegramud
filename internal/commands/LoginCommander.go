package commands

import (
    "gopkg.in/telegram-bot-api.v4"

    "github.com/Vehsamrak/telegramud/internal"
    "github.com/Vehsamrak/telegramud/internal/services"
)

type LoginCommander struct {
    AbstractCommander
    Stage    string
    Choice   ChoiceControl
    Database *services.Database
}

func (commander *LoginCommander) ExecuteCommand(command string, commandParameters []string) (
    commandResult internal.CommandResult,
) {
    var result string
    var executorName string

    if commander.Choice.Question != "" {
        commander.Choice.Answer = command

        if commander.Choice.CheckAnswer() {
            commander.Choice.ResultFunction(commander.Choice.Answer)
            commander.Stage = commander.Choice.AfterStage
            commander.Choice = ChoiceControl{}
        } else {
            result = commander.Choice.GetFailedMessage()
        }
    }

    if result == "" {
        switch commander.Stage {
        case "":
            commander.Choice = ChoiceControl{
                Question:         "Выбери свой класс",
                AvailableAnswers: []string{"воин", "лучник", "маг"},
                AfterStage:       "enterGame",
                ResultFunction: func(answer string) {
                    commander.connection.User.Race = commander.Choice.Answer
                    commander.connection.User.Save(commander.Database.GetConnection())
                },
            }

            result = commander.Choice.GetQuestionMessage()
        case "enterGame":
            result = "Ты осмотрелся.\n" + Look{}.Execute()
            executorName = "game"
        }
    }

    commandResult = internal.CommandResult{
        Message: tgbotapi.NewMessage(
            commander.connection.ChatId,
            result,
        ),
    }

    if executorName != "" {
        commandResult.ExecutorName = executorName
    }

    return commandResult
}
