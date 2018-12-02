package main

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Output struct {
	ChatID          int64
	Text            string
	message         tgbotapi.MessageConfig
	editMessage     tgbotapi.EditMessageTextConfig
	ReplyMarkup     interface{}
	callerMessageId int
}

func (output Output) SetMessage(message tgbotapi.MessageConfig) {
	output.message = message
}

func (output Output) SetReplyMarkup(markup interface{}) {
	output.ReplyMarkup = markup
}

func (output Output) SetText(text string) {
	output.Text = text
}

func (output Output) SetEditMessage(callerMessageId int, text string) {
	output.callerMessageId = callerMessageId
	output.Text = text
	// output.editMessage = tgbotapi.NewEditMessageText(output.ChatID, callerMessageId, text)
}

func (output Output) CreateChattable() tgbotapi.Chattable {
	var message tgbotapi.Chattable

	if output.editMessage.Text != "" {
		message = tgbotapi.EditMessageTextConfig{
			BaseEdit: tgbotapi.BaseEdit{
				ChatID:    output.ChatID,
				MessageID: output.callerMessageId,
			},
			Text: output.Text,
		}
	} else {
		message = &tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ReplyMarkup:      output.ReplyMarkup,
				ChatID:           output.ChatID,
				ReplyToMessageID: 0,
			},
			Text: output.Text,
			DisableWebPagePreview: false,
		}
	}

	return message
}
