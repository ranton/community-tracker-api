package peopleDetails

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	peopleRoutes := app.Group("/api/peopleDetails")
	peopleRoutes.Post("/", h.AddPeopleDetails)
	peopleRoutes.Put("/", h.UpdatePeopleDetails)
}
