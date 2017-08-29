package entities

import "github.com/jinzhu/gorm"

type Room struct {
    Id string `gorm:"primary_key;column:id"`
    Name string `gorm:"column:name"`
    Description string `gorm:"column:description"`
    Coordinates Coordinates
}

func (user *Room) Save(database *gorm.DB) {
    database.Save(user)
}
