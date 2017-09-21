package services

type WorldSaver struct {
    Database *Database
}

// Save all entities to database
func (sender *WorldSaver) Save() {
    sender.Database.GetConnection()
}
