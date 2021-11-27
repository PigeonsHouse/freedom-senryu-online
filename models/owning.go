package models

import "time"

type Owning struct {
	ID        string    `gorm:"id;primary_key"`
	RoomID    string    `gorm:"room_id"`
	UserID    string    `gorm:"user_id"`
	CreatedAt time.Time `gorm:"created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"updated_at;autoUpdateTime"`
}
