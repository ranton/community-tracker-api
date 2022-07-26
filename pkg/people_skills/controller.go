package people_skills

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
	route := app.Group("/api/peopleskills")
	route.Post("/", middleware.AuthMiddleware, h.AddPeopleSkills)
	route.Get("/", h.GetPeopleSkills)
	route.Put("/:peopleskillsid", middleware.AuthMiddleware, h.UpdateSkill)
	route.Delete("/:peopleskillsid", middleware.AuthMiddleware, h.DeleteSkills)
}
