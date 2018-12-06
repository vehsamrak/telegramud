package room

import "github.com/vehsamrak/telegramud/entity"

type RoomProvider struct {
	rooms []*Room
}

func (provider RoomProvider) Init() *RoomProvider {
	provider.rooms = []*Room{
		{id: "tavern", title: "Таверна \"Пьяный Докер\"",
			description: "Вы находитесь в шумной городской таверне.",
			actions: []*RoomAction{
				{commandName: "drink", commandTitle: "Заказать выпивку"},
				{commandName: "walk street", commandTitle: "Выйти на улицу"},
			},
		},
		{id: "street", title: "Портовая улица",
			description: "Портовая улица вымощена мраморными плитами, а по обеим сторонам возвышаются колонны. На улице многолюдно. Справа виднеется таверна \"Докер\".",
			actions: []*RoomAction{
				{commandName: "walk tavern", commandTitle: "Зайти в таверну"},
			},
		},
	}

	return &provider
}

func (provider *RoomProvider) FindById(roomName string) *Room {

	for _, room := range provider.rooms {
		if room.id == roomName {
			return room
		}
	}

	return nil
}

func (provider *RoomProvider) FindByPlayer(player *entity.Player) *Room {
	return provider.FindById(player.RoomId)
}
