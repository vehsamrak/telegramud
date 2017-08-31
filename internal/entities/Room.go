package entities

import (
    "github.com/jinzhu/gorm"
    "encoding/json"
)

type Room struct {
    Id string `gorm:"primary_key;column:id"`
    Name string `gorm:"column:name"`
    Description string `gorm:"column:description"`
    CoordinatesString string `gorm:"column:coordinates"`
    Coordinates *Coordinates
    Exits []*RoomPassage
}

func (room *Room) Save(database *gorm.DB) {
    database.Save(room)
}

func (room *Room) BeforeSave() (err error) {
    coordinatesJson, _ := json.Marshal(room.Coordinates)
    room.CoordinatesString = string(coordinatesJson)

    return
}
