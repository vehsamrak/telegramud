package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/vehsamrak/telegramud/entity"
)

const InitialRoomId = "tavern"

type PlayerProvider struct {
	players map[string]*entity.Player
}

func (provider PlayerProvider) Init() *PlayerProvider {
	provider.players = make(map[string]*entity.Player)

	return &provider
}

func (provider *PlayerProvider) FromTelegramUpdate(update tgbotapi.Update) *entity.Player {
	var chatId int64
	var username string

	if update.Message != nil {
		chatId = update.Message.Chat.ID
		username = update.Message.From.UserName
	} else if update.CallbackQuery != nil {
		chatId = update.CallbackQuery.Message.Chat.ID
		username = update.CallbackQuery.From.UserName
	}

	var player *entity.Player

	if provider.players[username] == nil {
		player = &entity.Player{ChatId: chatId, RoomId: InitialRoomId}
		provider.players[username] = player
	} else {
		player = provider.players[username]
	}

	return player
}

func (provider *PlayerProvider) AllPlayers() map[string]*entity.Player {
	return provider.players
}
