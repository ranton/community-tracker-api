package db

import (
	"fmt"
	"log"

	"github.com/VncntDzn/community-tracker-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	DB_HOST := config.GetEnv("DB_HOST")
	DB_PORT := config.GetEnv("DB_PORT")
	DB_NAME := config.GetEnv("DB_NAME")
	DB_USERNAME := config.GetEnv("DB_USERNAME")

	DB_PASSWORD := config.GetEnv("DB_PASSWORD")
	postgresConnection := fmt.Sprintf(
		"host=%s  port=%s   user=%s   password=%s  dbname=%s  sslmode=disable",
		DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_NAME)

	db, err := gorm.Open(postgres.Open(postgresConnection), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
