package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	_ "gorm.io/gorm"
)

func initDatabase() {
	connection := "user=vincentdizon dbname=mydb  host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
}
