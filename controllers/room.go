package controllers

import (
	"freedom-senryu-online/handlers"
	"freedom-senryu-online/middlewares"
	"freedom-senryu-online/models"

	"github.com/gin-gonic/gin"
)

type Room struct{}

var (
	hr handlers.Room
)

func (cr Room) CreateRoom(c *gin.Context) {
	var r models.Room

	uid := middlewares.GetCurrentUserInfo()

	r, err := hr.PostRoom(uid.UID)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(200, r)
}

func (cr Room) GetRoomByID(c *gin.Context) {
	var (
		r models.Room
	)

	rid := c.Params.ByName("room_id")
	r, err := hr.GetRoom(rid)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"detail": err.Error()})
		return
	}

	c.JSON(200, r)
}

func (cr Room) DeleteRoomByID(c *gin.Context) {
	uid := middlewares.GetCurrentUserInfo()

	rid := c.Params.ByName("room_id")
	err := hr.DeleteRoom(uid.UID, rid)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"detail": "OK",
	})
}
