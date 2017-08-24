package entities

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
    UserName string `gorm:"primary_key;column:name"`
    Race     string `gorm:"column:race"`
}

func (user *User) Save(database *gorm.DB) {
    database.Save(user)
}
