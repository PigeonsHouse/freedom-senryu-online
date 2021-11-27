package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"freedom-senryu-online/models"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	err = godotenv.Load(".env")
	var (
		p_user     string
		p_db       string
		p_password string
	)
	if err == nil {
		p_user = os.Getenv("POSTGRES_USER")
		p_db = os.Getenv("POSTGRES_DB")
		p_password = os.Getenv("POSTGRES_PASSWORD")
	} else {
		p_user = "gorm"
		p_db = "gorm"
		p_password = "gorm"
	}
	db_opning_str := fmt.Sprintf("host=127.0.0.1 port=15432 user=%s dbname=%s password=%s sslmode=disable", p_user, p_db, p_password)
	db, err = gorm.Open("postgres", db_opning_str)
	if err != nil {
		panic(err)
	}

	autoMigration()
}

func GetDB() *gorm.DB {
	return db
}

func Clone() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.Senryu{},
		&models.Favorite{},
	)
}
