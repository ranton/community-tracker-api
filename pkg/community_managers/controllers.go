package community_managers

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
	routes := app.Group("/api/managers")
	routes.Get("/", h.GetManagers)
	routes.Get("/community", h.GetCommunityManagers)
}
