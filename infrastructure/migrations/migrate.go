package main

import (
    "os"
    "fmt"

    "database/sql"
    "github.com/rubenv/sql-migrate"
    _ "github.com/lib/pq"
)

const DATABASE_DRIVER = "postgres"
const DATABASE_NAME = "telegramud"
const DATABASE_USER = "telegramud"
const DATABASE_HOST = "database:5432/"
const DATABASE_DATA_SOURCE = "postgres://" + DATABASE_HOST + DATABASE_NAME + "?sslmode=disable;user=" + DATABASE_USER
const MIGRATIONS_PATH = "infrastructure/migrations"

// Run database migrations
func main() {
    fmt.Print("Migrating database...")

    migrations := &migrate.FileMigrationSource{Dir: MIGRATIONS_PATH}

    database, error := sql.Open(DATABASE_DRIVER, DATABASE_DATA_SOURCE)
    defer database.Close()

    if error != nil {
        fmt.Println(error)
    }

    executedMigrationsCount, error := migrate.Exec(database, DATABASE_DRIVER, migrations, migrate.Up)
    if error != nil {
        fmt.Println(error)
    }

    fmt.Printf("done. Applied %d migrations!\n", executedMigrationsCount)

    os.Exit(0)
}
