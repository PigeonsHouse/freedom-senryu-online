package main

import (
	"freedom-senryu-online/db"
	"freedom-senryu-online/middlewares"
	"freedom-senryu-online/routes"
)

func main() {
	middlewares.InitClient()
	db.Init()
	routes.Init()
}
