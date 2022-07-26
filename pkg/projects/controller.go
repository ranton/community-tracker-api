package projects

import (
	"github.com/gofiber/fiber/v2"
	"github.com/VncntDzn/community-tracker-api/pkg/middleware"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	projectRoutes := app.Group("/api/projects")
	projectRoutes.Post("/", middleware.AuthMiddleware, h.CreateProject)
	projectRoutes.Get("/:projectid", h.GetProject)
}
