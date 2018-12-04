package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Output struct {
	ChatID               int64
	Text                 string
	editMessage          tgbotapi.EditMessageTextConfig
	ReplyMarkup          interface{}
	InlineKeyboardMarkup *tgbotapi.InlineKeyboardMarkup
	callerMessageId      int
	callerMessageText    string
	isEdit               bool
}

func (output *Output) SetReplyMarkup(markup interface{}) {
	output.ReplyMarkup = markup
}

func (output *Output) SetText(text string) {
	output.Text = text
}

func (output *Output) SetEditText(text string) {
	output.isEdit = true
	output.Text = text
}

func (output *Output) SetEditKeyboard(markup *tgbotapi.InlineKeyboardMarkup) {
	output.isEdit = true
	output.InlineKeyboardMarkup = markup
}

func (output *Output) GenerateChattable() tgbotapi.Chattable {
	var message tgbotapi.Chattable

	if output.isEdit {
		if output.Text == "" {
			output.Text = output.callerMessageText
		}

		message = tgbotapi.EditMessageTextConfig{
			BaseEdit: tgbotapi.BaseEdit{
				ChatID:      output.ChatID,
				ReplyMarkup: output.InlineKeyboardMarkup,
				MessageID:   output.callerMessageId,
			},
			Text:      output.Text,
			ParseMode: "markdown",
		}
	} else {
		message = &tgbotapi.MessageConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID:           output.ChatID,
				ReplyMarkup:      output.ReplyMarkup,
				ReplyToMessageID: 0,
			},
			Text: output.Text,
			DisableWebPagePreview: false,
			ParseMode:             "markdown",
		}
	}

	return message
}
