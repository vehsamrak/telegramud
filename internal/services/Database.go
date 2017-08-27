package services

import "github.com/jinzhu/gorm"

type Database struct {
    connection *gorm.DB
}

func (database *Database) GetConnection() *gorm.DB {
    if database.connection != nil {
        return database.connection
    }

    database.connection, _ = gorm.Open(
        "postgres",
        "host=localhost port=5432 user=telegramud password=telegramud dbname=telegramud sslmode=disable",
    )

    return database.connection
}
