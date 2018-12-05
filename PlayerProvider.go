package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type PlayerProvider struct {
	players map[string]*Player
}

func (provider PlayerProvider) Init() *PlayerProvider {
	provider.players = make(map[string]*Player)

	return &provider
}

func (provider *PlayerProvider) FromTelegramUpdate(update tgbotapi.Update) *Player {
	var chatId int64
	var username string

	if update.Message != nil {
		chatId = update.Message.Chat.ID
		username = update.Message.From.UserName
	} else if update.CallbackQuery != nil {
		chatId = update.CallbackQuery.Message.Chat.ID
		username = update.CallbackQuery.From.UserName
	}

	var player *Player

	if provider.players[username] == nil {
		player = &Player{ChatId: chatId}
		provider.players[username] = player
	} else {
		player = provider.players[username]
	}

	return player
}

func (provider *PlayerProvider) AllPlayers() map[string]*Player {
	return provider.players
}
