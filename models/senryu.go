package models

import (
	"time"
)

type Senryu struct {
	ID             string    `json:"id"`
	Theme          string    `json:"theme"`
	FirstLength    int       `json:"first_length"`
	FirstContent   string    `json:"first_content"`
	FirstAuthorID  string    `json:"first_author_id"`
	SecondLength   int       `json:"second_length"`
	SecondContent  string    `json:"second_content"`
	SecondAuthorID string    `json:"second_author_id"`
	ThirdLength    int       `json:"third_length"`
	ThirdContent   string    `json:"third_content"`
	ThirdAuthorID  string    `json:"third_author_id"`
	RoomID         string    `json:"room_id"`
	IsFinished     bool      `json:"is_finished"`
	IsPublic       bool      `json:"is_public"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
