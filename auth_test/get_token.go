package main

import (
	"context"
	"fmt"
	"freedom-senryu-online/middlewares"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("auth_test/.env")
	if err != nil {
		fmt.Println("cannot find env file")
		return
	}

	api_key := os.Getenv("GOOGLE_API_KEY")

	middlewares.InitClient()
	client := middlewares.GetClient()
	token, err := client.CustomToken(context.Background(), api_key)
	if err == nil {
		fmt.Println(token)
	}
}
