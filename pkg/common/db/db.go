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

	db.AutoMigrate(&models.Cities{})

	return db
}

/* dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
// Connect to the DB and initialize the DB variable
DB, err = gorm.Open(postgres.Open(dsn))

if err != nil {
	panic("failed to connect database")
}

fmt.Println("Connection Opened to Database") */
