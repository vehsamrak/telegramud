package entities

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
    UserName string `gorm:"primary_key;column:name"`
    Race     string `gorm:"column:race"`
    RoomId   string `gorm:"column:room;type:varchar(100)"`
    Room     *Room  `gorm:"ForeignKey:RoomId"`
    ChatId   int64  `gorm:"column:chat_id"`
    Executor string `gorm:"column:executor;type:varchar(100)"`
}

func (user *User) Save(database *gorm.DB) {
    database.Save(user)
}

func (user *User) FindByName(database *gorm.DB, username string) *User {
    database.Where("name = ?", username).First(user)

    return user
}

func (user *User) FindAll(database *gorm.DB) []User {

    users := []User{}

    database.Find(&users)

    return users
}
