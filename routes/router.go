package routes

import (
	"freedom-senryu-online/controllers"
	"freedom-senryu-online/middlewares"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	ctrl_user controllers.User
	ctrl_room controllers.Room
	ctrl_ws   controllers.WebSocket
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Init() {
	r := gin.Default()
	cors := cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			"POST",
			"GET",
			"PATCH",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
	})
	r.Use(cors)

	g := r.Group("/api/v1")
	g.Use(middlewares.FirebaseAuthorization())

	UserRoutes(g)
	RoomRoutes(g)
	SenryuRoutes(g)
	FavoriteRoutes(g)
	WebSocketRoutes(g)

	r.Run(":8080")
}

func UserRoutes(g *gin.RouterGroup) {
	u := g.Group("/users")
	u.POST("", ctrl_user.CreateUser)
	u.PUT("", ctrl_user.PutMe)
	u.DELETE("", ctrl_user.DeleteMe)
	u.GET("/:user_id", ctrl_user.GetUser)
	u.GET("/me", ctrl_user.GetMe)
	u.GET("/:user_id/senryus", func(c *gin.Context) {})
	u.GET("/me/senryus", func(c *gin.Context) {})
}

func RoomRoutes(g *gin.RouterGroup) {
	r := g.Group("/rooms")
	r.POST("", ctrl_room.CreateRoom)
	r.GET("/:room_id", ctrl_room.GetRoomByID)
	r.DELETE("/:room_id", ctrl_room.DeleteRoomByID)
}

func SenryuRoutes(g *gin.RouterGroup) {
	s := g.Group("/senryus")
	s.POST("", func(c *gin.Context) {})
	s.GET("", func(c *gin.Context) {})
	s.PUT("/:senryu_id", func(c *gin.Context) {})
	s.GET("/:senryu_id", func(c *gin.Context) {})
	s.DELETE("/:senryu_id", func(c *gin.Context) {})
}

func FavoriteRoutes(g *gin.RouterGroup) {
	f := g.Group("/favorites")
	f.POST("/:senryu_id", func(c *gin.Context) {})
	f.DELETE("/:senryu_id", func(c *gin.Context) {})
}

func WebSocketRoutes(g *gin.RouterGroup) {
	f := g.Group("/ws")
	f.GET("/:room_id", ctrl_ws.ControllWebSocket)
}
