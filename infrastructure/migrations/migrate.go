package main

import (
    "os"
    "fmt"

    "github.com/vehsamrak/telegramud/internal/services"
    "github.com/vehsamrak/telegramud/internal/entities"
)

// Run automatic migrations on database if flag -migrate passed
func main() {
    database := &services.Database{}
    databaseConnection := database.GetConnection()
    defer databaseConnection.Close()

    fmt.Print("Migrating database...")

    databaseConnection.AutoMigrate(&entities.User{}, &entities.Room{})

    fmt.Println("done.")

    os.Exit(0)
}
