package main

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/db"
	"github.com/VncntDzn/community-tracker-api/pkg/community"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	myDB := db.Init()

	community.RegisterRoutes(app, myDB)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString(":8000")
	})

	// books.RegisterRoutes(app, db)

	app.Listen(":8000")
}
