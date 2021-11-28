package controllers

import (
	"freedom-senryu-online/handlers"
	"freedom-senryu-online/middlewares"
	"freedom-senryu-online/models"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserInfo **auth.UserInfo
}

var (
	hu handlers.User
)

func (cu User) CreateUser(c *gin.Context) {
	var u models.User
	hu = handlers.User{}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.AbortWithStatusJSON(422, gin.H{"detail": err.Error()})
		return
	}

	ui := middlewares.GetCurrentUserInfo()

	u, err := hu.PostUser(ui.UID, u.Email, u.Name)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(200, u)
}

func (cu User) GetUser(c *gin.Context) {
	uid := c.Params.ByName("user_id")
	u, err := hu.GetByID(uid)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(200, u)
}

func (cu User) GetMe(c *gin.Context) {
	ui := middlewares.GetCurrentUserInfo()

	u, err := hu.GetByID(ui.UID)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(200, u)
}

func (cu User) PutMe(c *gin.Context) {
	name := c.Query("name")

	ui := middlewares.GetCurrentUserInfo()

	u, err := hu.UpdateUserName(ui.UID, name)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(200, u)
}

func (cu User) DeleteMe(c *gin.Context) {
	ui := middlewares.GetCurrentUserInfo()

	err := hu.DeleteByID(ui.UID)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"detail": "OK",
	})
}
