package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var PortApp string

func LoadConfig() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("The \"config.env\" file is invalid (please rename 'config.example.env' to 'config.env'")
	}

	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
	dsn = fmt.Sprintf(dsn, os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"), os.Getenv("DB_TIMEZONE"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	PortApp = os.Getenv("PORT_APP")
}
