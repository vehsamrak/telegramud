package services

import (
    "github.com/jinzhu/gorm"

    "github.com/Vehsamrak/telegramud/internal/entities"
)

type RoomGenerator struct{}

func (generator *RoomGenerator) GenerateWorld(database *gorm.DB) []*entities.Room {
    market := &entities.Room{
        Id:          "market",
        Name:        "Рыночная площадь",
        Description: "Площадь с шатрами и лавками торговцев.",
        Coordinates: &entities.Coordinates{X: 1, Y: 2, Z: 0},
    }

    tavern := &entities.Room{
        Id:          "tavern",
        Name:        "Таверна",
        Description: "Шумная городская таверна, где начинаются приключения.",
        Coordinates: &entities.Coordinates{X: 1, Y: 1, Z: 0},
    }

    street := &entities.Room{
        Id:          "street",
        Name:        "Улица",
        Description: "Шумная городская улица.",
        Coordinates: &entities.Coordinates{X: 1, Y: 3, Z: 0},
    }

    market.Passages = append(market.Passages, &entities.RoomPassage{RoomTo: tavern})
    tavern.Passages = append(tavern.Passages, &entities.RoomPassage{RoomTo: market})
    tavern.Passages = append(tavern.Passages, &entities.RoomPassage{RoomTo: street})

    rooms := []*entities.Room{
        tavern,
        market,
    }

    for _, room := range rooms {
        room.Save(database)
    }

    return rooms
}
