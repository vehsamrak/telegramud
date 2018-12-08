package room

import "github.com/vehsamrak/telegramud/entity"

type RoomProvider struct {
	rooms map[string]*Room
}

func (provider RoomProvider) Init() *RoomProvider {
	rooms := []*Room{
		{id: "tavern", title: "Таверна \"Пьяный Докер\"",
			description: "Вы находитесь в шумной городской таверне.",
			actions: []*RoomAction{
				{actionName: "Заказать выпивку", command: "drink"},
			},
			exits: []*RoomExit{
				{destinationRoomId: "street"},
			},
		},
		{id: "street", title: "Портовая улица",
			description: "Портовая улица вымощена мраморными плитами, а по обеим сторонам возвышаются колонны. На улице многолюдно. Справа виднеется таверна \"Докер\".",
			exits: []*RoomExit{
				{destinationRoomId: "port"},
				{destinationRoomId: "town_gate"},
				{destinationRoomId: "tavern"},
			},
		},
		{id: "port", title: "Порт",
			description: "Порт находится на юге города. Большие торговые суда и военные корабли осуществляют погрузку. Некрупные прогулочные яхты неспешно качаются на волнах, ударяюсь бортами о пирс.",
			exits: []*RoomExit{
				{destinationRoomId: "street"},
			},
		},
		{id: "town_gate", title: "Городские ворота",
			description: "Массивные деревянные ворота надежно защищают город от окружающего леса.",
			exits: []*RoomExit{
				{destinationRoomId: "road1"},
				{destinationRoomId: "street"},
			},
		},
		{id: "road1", title: "Дорога",
			description: "За городскими воротами начинается дорога, ведущая в лес. Вдали виднеются горы, окруженные высокими деревьями.",
			exits: []*RoomExit{
				{destinationRoomId: "forest_road1"},
				{destinationRoomId: "town_gate"},
			},
		},
		{id: "forest_road1", title: "Лесная дорога",
			description: "Лес по сторонам сгущается, нависая кронами деревьев, но мощеная дорога не сужается. Монумент местных каменщиков.",
			exits: []*RoomExit{
				{destinationRoomId: "forest1"},
				{destinationRoomId: "road1"},
			},
		},
		{id: "forest1", title: "Лесная опушка",
			description: "Лес, окончательно победив дорогу начинает здесь входить в полную силу. На опушке виднеется пара бревенчатых срубов и множество пеньков после лесозаготовки.",
			exits: []*RoomExit{
				{destinationRoomId: "forest2"},
				{destinationRoomId: "forest_road1"},
			},
		},
		{id: "forest2", title: "Полесье",
			description: "Кроны раскидистых деревев пропускают солнечный свет на зеленую траву.",
			exits: []*RoomExit{
				{destinationRoomId: "forest3"},
				{destinationRoomId: "forest1"},
			},
		},
		{id: "forest3", title: "Лес",
			description: "Лиственный лес с дубами и березами. Слышно веселое пение птиц.",
			exits: []*RoomExit{
				{destinationRoomId: "forest2"},
				{destinationRoomId: "forest_river1"},
				{destinationRoomId: "forest4"},
			},
		},
		{id: "forest_river1", title: "Лесная речушка",
			description: "Порожистая речка средних размеров пересекает лес. В воде, разбрасывая брызги, плещутся рыбы.",
			exits: []*RoomExit{
				{destinationRoomId: "forest4"},
				{destinationRoomId: "forest3"},
			},
		},
		{id: "forest4", title: "Чащоба",
			description: "Вы забрались глубоко в лесную чащобу. Кроны деревьев закрывают небо, не давая проходить свету и осадкам.",
			exits: []*RoomExit{
				{destinationRoomId: "forest3"},
			},
		},
	}

	provider.rooms = make(map[string]*Room)

	for _, room := range rooms {
		provider.rooms[room.Id()] = room
	}

	return &provider
}

func (provider *RoomProvider) FindById(roomName string) *Room {
	return provider.rooms[roomName]
}

func (provider *RoomProvider) FindByPlayer(player *entity.Player) *Room {
	return provider.FindById(player.RoomId)
}

func (provider *RoomProvider) FindAll() map[string]*Room {
	return provider.rooms
}
