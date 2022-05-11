package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()

	connection := "user=vincentdizon dbname=mydb  host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")

	app.Get("/", func(c *fiber.Ctx) error {
		res := (db.QueryRow(`SELECT * FROM "PEOPLE"`))
		return c.JSONP(res)
	})

	app.Listen(":3000")
}
