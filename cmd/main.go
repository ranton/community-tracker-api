package main

import (
	"github.com/VncntDzn/community-tracker-api/pkg/books"
	"github.com/VncntDzn/community-tracker-api/pkg/common/db"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	sad := db.Init()

	books.RegisterRoutes(app, sad)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString(":8000")
	})

	// books.RegisterRoutes(app, db)

	app.Listen(":8000")
}
