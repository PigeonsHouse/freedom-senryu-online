package handlers

import (
	"errors"
	"freedom-senryu-online/db"
	"freedom-senryu-online/models"
)

type Room struct{}

func (hr Room) PostRoom(uid string) (models.Room, error) {
	db := db.GetDB()
	var (
		user models.User
		hu   User
		room models.Room
		err  error
	)
	if user, err = hu.GetByID(uid); err != nil {
		return room, err
	}
	if user.RoomID != "" {
		return room, errors.New("this User already belong at the room")
	}
	room = models.Room{
		IsPlaying:   false,
		Owner:       []models.User{user},
		JoiningUser: []models.User{user},
	}
	if err := db.Create(&room).Error; err != nil {
		return room, err
	}

	db.Model(&user).Association("Room").Append(&room)

	return room, nil
}

func (hr Room) GetRoom(uid string, rid string) (models.Room, error) {
	db := db.GetDB()

	var (
		room         models.Room
		owner        []models.User
		joining_user []models.User
	)
	if err := db.First(&room, "id = ?", rid).Error; err != nil {
		return room, err
	}
	if err := db.Model(&room).Association("Owner").Find(&owner).Error; err != nil {
		return room, err
	}
	if err := db.Model(&room).Association("Owner").Find(&joining_user).Error; err != nil {
		return room, err
	}

	return room, nil
}

// func (hr Room) DeleteRoom() error
