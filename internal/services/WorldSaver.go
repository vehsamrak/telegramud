package services

import "fmt"

type WorldSaver struct {
    Database *Database
}

// Save all entities to database
func (sender *WorldSaver) Save() {
    sender.Database.GetConnection().Close()
    fmt.Println("Here comes the world saving ...")
}
