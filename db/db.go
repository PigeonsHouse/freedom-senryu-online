package db

import (
	"fmt"
	"os"
	"strconv"

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
	p_user := "gorm"
	p_db := "gorm"
	p_password := "gorm"
	p_host := "localhost"
	p_port := 5432
	if err == nil {
		p_user = os.Getenv("POSTGRES_USER")
		p_db = os.Getenv("POSTGRES_DB")
		p_password = os.Getenv("POSTGRES_PASSWORD")
		p_host = os.Getenv("POSTGRES_HOST")
		p_port, err = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	}
	db_opning_str := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", p_host, p_port, p_user, p_db, p_password)
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
