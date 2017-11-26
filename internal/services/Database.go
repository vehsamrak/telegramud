package services

import "github.com/jinzhu/gorm"

type Database struct {
    connection *gorm.DB
    Password   string
}

func (database *Database) GetConnection() *gorm.DB {
    if database.connection != nil {
        return database.connection
    }

    database.connection, _ = gorm.Open(
        "postgres",
        "host=telegramud_database port=5432 user=telegramud password=" + database.Password + " dbname=telegramud sslmode=disable",
    )

    return database.connection
}
