package services

import "fmt"

type WorldSaver struct {
    Database *Database
    Players  map[string]*Connection
}

// Save all entities to database
func (world *WorldSaver) Save() {
    world.Database.GetConnection().Close()
    fmt.Println("Here comes the world saving ...")
}

// Initialise world saver
func (world *WorldSaver) Init() {
    world.Players = map[string]*Connection{}
}
