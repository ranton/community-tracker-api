package admin

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
	route := app.Group("/api/admin")
	route.Post("/", middleware.AuthMiddleware, h.CreateAdmin)
	route.Put("/:communityadminandmanagerid", middleware.AuthMiddleware, h.UpdateAdminDetails)
	route.Put("/:communityadminandmanagerid/password", middleware.AuthMiddleware, h.UpdatePassword)
}
