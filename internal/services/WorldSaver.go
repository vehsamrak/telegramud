package services

import (
    "fmt"

    "github.com/vehsamrak/telegramud/internal/entities"
)

type WorldSaver struct {
    Database *Database
    Players  map[string]*Connection
}

// SaveWorld all entities to database
func (world *WorldSaver) SaveWorld() {
    fmt.Println("Here comes the world saving ...")

    for _, connection := range world.Players {
        connection.User.Save(world.Database.GetConnection())
        fmt.Printf("Saving user %v.\n", connection.User.UserName)
    }
}

// Initialise world saver
func (world *WorldSaver) Init() {
    world.Players = map[string]*Connection{}
    userService := &entities.User{}
    allUsers := userService.FindAll(world.Database.GetConnection())

    for _, user := range allUsers {
        world.Players[user.UserName] = &Connection{
            User:   &user,
        }
    }

}

// Destruct world
func (world *WorldSaver) Destruct() {
    world.Database.GetConnection().Close()
}
