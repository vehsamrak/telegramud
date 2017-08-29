package services

import "github.com/Vehsamrak/telegramud/internal/entities"

type RoomGenerator struct{}

func (generator *RoomGenerator) GenerateWorld() []*entities.Room {
    rooms := []*entities.Room{
        {
            Id:          "tavern",
            Name:        "Таверна",
            Description: "Шумная городская таверна, где начинаются приключения.",
            Coordinates: &entities.Coordinates{X: 1, Y: 1, Z: 0},
        },
    }

    return rooms
}
