package db

import (
	"log"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=vincentdizon dbname=mydb sslmode=disable"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Community{})

	return db
}
