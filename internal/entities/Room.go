package entities

import (
    "github.com/jinzhu/gorm"
    "encoding/json"
)

type Room struct {
    Id                string `gorm:"primary_key;column:id;type:varchar(100)"`
    Name              string `gorm:"column:name;type:varchar(100)"`
    Description       string `gorm:"column:description"`
    CoordinatesString string `gorm:"column:coordinates;unique_index"`
    Coordinates       *Coordinates
    Passages          []*RoomPassage
}

func (room *Room) Save(database *gorm.DB) {
    database.Save(room)
}

func (room *Room) BeforeSave() (err error) {
    coordinatesJson, _ := json.Marshal(room.Coordinates)
    room.CoordinatesString = string(coordinatesJson)

    return
}

func (room *Room) AfterFind() (err error) {
    json.Unmarshal([]byte(room.CoordinatesString), room.Coordinates)

    return
}
