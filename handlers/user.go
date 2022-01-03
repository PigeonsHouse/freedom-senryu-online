package handlers

import (
	"freedom-senryu-online/db"
	"freedom-senryu-online/models"

	"github.com/go-playground/validator/v10"
)

type User struct{}

func (hu User) PostUser(uid string, email string, name string) (models.User, error) {
	validate := validator.New()
	db := db.GetDB()
	user := models.User{ID: uid, Email: email, Name: name}
	if err := validate.Struct(user); err != nil {
		return user, err
	}
	if err := db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (hu User) GetByID(uid string) (models.User, error) {
	db := db.GetDB()
	var u models.User
	if err := db.Where("id = ?", uid).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (hu User) UpdateUserName(uid string, name string) (models.User, error) {
	db := db.GetDB()
	var (
		u   models.User
		err error
	)
	if err := db.Model(&u).Where("id = ?", uid).Update("name", name).Error; err != nil {
		return u, nil
	}
	if u, err = hu.GetByID(uid); err != nil {
		return u, nil
	}
	return u, nil
}

func (hu User) DeleteByID(uid string) error {
	db := db.GetDB()
	if err := db.Where("id = ?", uid).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (hu User) JoinRoom(uid string, rid string) (models.User, error) {
	var (
		u   models.User
		err error
	)
	db := db.GetDB()
	if err = db.Model(&u).Where("id = ?", uid).Update("room_id", rid).Error; err != nil {
		return u, err
	}
	if u, err = hu.GetByID(uid); err != nil {
		return u, nil
	}
	return u, nil
}

func (hu User) ExitRoom(uid string) error {
	db := db.GetDB()
	var u models.User
	if err := db.Model(&u).Where("id = ?", uid).Update("room_id", nil).Error; err != nil {
		return err
	}
	return nil
}
