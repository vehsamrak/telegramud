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
				{commandName: "Заказать выпивку"},
				{commandName: "Выйти на улицу"},
			},
		},
		{id: "street", title: "Портовая улица",
			description: "Портовая улица вымощена мраморными плитами, а по обеим сторонам возвышаются колонны. На улице многолюдно. Справа виднеется таверна \"Докер\".",
			actions: []*RoomAction{
				{commandName: "Порт"},
				{commandName: "Городские ворота"},
				{commandName: "Зайти в таверну"},
			},
		},
		{id: "port", title: "Порт",
			description: "Порт находится на юге города. Большие торговые суда и военные корабли осуществляют погрузку. Некрупные прогулочные яхты неспешно качаются на волнах, ударяюсь бортами о пирс.",
			actions: []*RoomAction{
				{commandName: "Портовая улица"},
			},
		},
		{id: "town_gate", title: "Городские ворота",
			description: "Массивные деревянные ворота надежно защищают город от окружающего леса.",
			actions: []*RoomAction{
				{commandName: "Дорога"},
				{commandName: "Портовая улица"},
			},
		},
		{id: "road1", title: "Дорога",
			description: "За городскими воротами начинается дорога, ведущая в лес. Вдали виднеются горы, окруженные высокими деревьями.",
			actions: []*RoomAction{
				{commandName: "Лесная дорога"},
				{commandName: "Городские ворота"},
			},
		},
		{id: "forest_road1", title: "Лесная дорога",
			description: "Лес по сторонам сгущается, нависая кронами деревьев, но мощеная дорога не сужается. Монумент местных каменщиков.",
			actions: []*RoomAction{
				{commandName: "Лесная опушка"},
				{commandName: "Дорога"},
			},
		},
		{id: "forest1", title: "Лесная опушка",
			description: "Лес, окончательно победив дорогу начинает здесь входить в полную силу. На опушке виднеется пара бревенчатых срубов и множество пеньков после лесозаготовки.",
			actions: []*RoomAction{
				{commandName: "Полесье"},
				{commandName: "Лесная дорога"},
			},
		},
		{id: "forest2", title: "Полесье",
			description: "Кроны раскидистых деревев пропускают солнечный свет на зеленую траву.",
			actions: []*RoomAction{
				{commandName: "Лес"},
				{commandName: "Лесная опушка"},
			},
		},
		{id: "forest3", title: "Лес",
			description: "Лиственный лес с дубами и березами. Слышно веселое пение птиц.",
			actions: []*RoomAction{
				{commandName: "Полесье"},
				{commandName: "Лесная речушка"},
				{commandName: "Чащоба"},
			},
		},
		{id: "forest_river1", title: "Лесная речушка",
			description: "Порожистая речка средних размеров пересекает лес. В воде, разбрасывая брызги, плещутся рыбы.",
			actions: []*RoomAction{
				{commandName: "Чащоба"},
				{commandName: "Лес"},
			},
		},
		{id: "forest4", title: "Чащоба",
			description: "Вы забрались глубоко в лесную чащобу. Кроны деревьев закрывают небо, не давая проходить свету и осадкам.",
			actions: []*RoomAction{
				{commandName: "Лес"},
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
