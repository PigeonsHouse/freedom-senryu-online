package models

import (
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" validate:"email"`
	Name      string    `json:"name"`
	RoomID    *string   `json:"living_room_id"`
	IsOwner   bool      `json:"is_owner"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
