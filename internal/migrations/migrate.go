package migrations

import (
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "github.com/Vehsamrak/telegramud/internal/services"
    "github.com/Vehsamrak/telegramud/internal/entities"
    "fmt"
)

func Migrate() {
    database := &services.Database{}
    databaseConnection := database.GetConnection()
    defer databaseConnection.Close()

    fmt.Println("Migrating database...")

    databaseConnection.AutoMigrate(&entities.User{}, &entities.Room{})
}
