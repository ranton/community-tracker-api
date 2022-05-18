package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	_ "gorm.io/gorm"
)

type people struct {
	PEOPLE_ID  string `json:"people_id"`
	FIRST_NAME string `json:"first_name"`
}

func setUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi")
	})
}
func main() {
	app := fiber.New()
	setUpRoutes(app)

	app.Listen(":3000")

}
