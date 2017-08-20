package commands

import (
	"fmt"
	"strings"
)

type ChoiceControl struct {
	Question         string
	AvailableAnswers []string
	Answer           string
	AfterStage       string
}

func (choice *ChoiceControl) GetFailedMessage() string {
	return fmt.Sprintf("Неверный выбор. Доступно: *%v*.", strings.Join(choice.AvailableAnswers, "*, *"))
}

func (choice *ChoiceControl) GetQuestionMessage() string {
	return fmt.Sprintf("%v: *%v*.", choice.Question, strings.Join(choice.AvailableAnswers, "*, *"))
}

func (choice *ChoiceControl) CheckAnswer() bool {
	return contains(choice.Answer, choice.AvailableAnswers)
}

func contains(e string, s []string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
