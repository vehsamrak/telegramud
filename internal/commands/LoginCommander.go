package commands

import (
    "gopkg.in/telegram-bot-api.v4"

    "github.com/Vehsamrak/telegramud/internal"
)

type LoginCommander struct {
    AbstractCommander
    Stage  string
    Choice ChoiceControl
}

func (commander *LoginCommander) ExecuteCommand(command string, commandParameters []string) (
    commandResult internal.CommandResult,
) {
    var result string
    var executorName string

    if commander.Choice.Question != "" {
        commander.Choice.Answer = command

        if commander.Choice.CheckAnswer() {
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
