package cities

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
	citiesRoutes := app.Group("/api/cities")
	citiesRoutes.Get("/", h.GetCities)
	citiesRoutes.Post("/", h.AddCity)
	citiesRoutes.Delete("/:id", h.DeleteCity)
	citiesRoutes.Put("/:id", h.UpdateCity)
}
