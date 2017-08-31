package services

import (
    "github.com/Vehsamrak/telegramud/internal/entities"
    "github.com/jinzhu/gorm"
)

type RoomGenerator struct{}

func (generator *RoomGenerator) GenerateWorld(database *gorm.DB) []*entities.Room {
    rooms := []*entities.Room{
        {
            Id:          "tavern",
            Name:        "Таверна",
            Description: "Шумная городская таверна, где начинаются приключения.",
            Coordinates: &entities.Coordinates{X: 1, Y: 1, Z: 0},
        },
        {
            Id:          "marketplace",
            Name:        "Рыночная площадь",
            Description: "Площадь с шатрами и лавками торговцев.",
            Coordinates: &entities.Coordinates{X: 1, Y: 2, Z: 0},
        },
    }

    for _, room := range rooms {
        room.Save(database)
    }

    return rooms
}
