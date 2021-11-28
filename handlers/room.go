package handlers

import (
	"errors"
	"fmt"
	"freedom-senryu-online/db"
	"freedom-senryu-online/models"

	"github.com/jinzhu/gorm"
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
	if user.RoomID != nil {
		return room, errors.New("this User already belong at the room")
	}
	if err := db.Model(&user).Where("id = ?", uid).Update("is_owner", true).Error; err != nil {
		return room, nil
	}

	room = models.Room{
		IsPlaying: false,
		Owner:     user,
		Users:     []models.User{user},
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
		joining_user []models.User
	)
	if err := db.First(&room, "id = ?", rid).Error; err != nil {
		return room, err
	}
	if err := db.Model(&room).Association("Users").Find(&joining_user).Error; err != nil {
		return room, err
	}
	owner := models.User{}

	for _, user := range joining_user {
		if user.IsOwner {
			owner = user
			break
		}
	}
	room.Users = joining_user
	room.Owner = owner

	return room, nil
}

func (hr Room) DeleteRoom(uid string, rid string) error {
	db := db.GetDB()

	deleting_user := models.User{}

	if err := db.Where("id = ?", uid).First(&deleting_user); err != nil {
		if *deleting_user.RoomID != rid || !deleting_user.IsOwner {
			return errors.New("you can not delete room you don't owning")
		}
	}

	room := models.Room{}
	if err := db.First(&room, "id = ?", rid).Error; err != nil {
		return err
	}
	joining_user := []models.User{}
	if err := db.Model(&room).Association("Users").Find(&joining_user).Error; err != nil {
		return err
	}

	for _, user := range joining_user {
		fmt.Println(user.ID)
		if err := db.Model(&user).Update("is_owner", false).Error; err != nil {
			return err
		}
		if err := db.Model(&user).Update("room_id", gorm.Expr("NULL")).Error; err != nil {
			return err
		}
	}

	if err := db.Where("id = ?", rid).Delete(&models.Room{}).Error; err != nil {
		return err
	}
	return nil
}
