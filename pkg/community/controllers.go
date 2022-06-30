package community

import (
	"github.com/VncntDzn/community-tracker-api/pkg/middleware"
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
	communityRoutes := app.Group("/api/community")
	communityRoutes.Get("/job-level", h.GetJobLevel)
	communityRoutes.Get("/", h.GetCommunity)
	communityRoutes.Post("/", middleware.AuthMiddleware, h.AddCommunity)
	communityRoutes.Put("/:communityid", middleware.AuthMiddleware, h.UpdateCommunity)
}
