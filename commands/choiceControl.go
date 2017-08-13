package commands

import (
    "fmt"
    "strings"
)

type ChoiceControl struct {
    Question         string
    AvailableAnswers map [string] bool
    Answer           string
}

func (choice *ChoiceControl) GetFailedMessage() string {
    return fmt.Sprint("Неверный выбор.")
}

func (choice *ChoiceControl) GetQuestionMessage() string {
    return fmt.Sprintf("%v. [%v]", choice.Question, strings.Join(choice.getAvailableAnswers(), ", "))
}

func (choice *ChoiceControl) CheckAnswer() bool {
    return choice.AvailableAnswers[choice.Answer]
}

func (choice *ChoiceControl) getAvailableAnswers() []string {
    var readableAnswers []string

    for answer := range choice.AvailableAnswers {
        readableAnswers = append(readableAnswers, answer)
    }

    return readableAnswers
}
