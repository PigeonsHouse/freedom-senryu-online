package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Room struct {
	ID        string    `json:"id" gorm:"primary_key"`
	IsPlaying bool      `json:"is_playing"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Owner     User      `json:"owner_user"`
	Users     []User    `json:"users" gorm:"foreignKey:RoomID"`
}

func (room *Room) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid.String())
}
