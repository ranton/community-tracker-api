package people_skills

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
	route := app.Group("/api/peopleskills")
	route.Post("/", h.AddPeopleSkills)
	route.Get("/", h.GetPeopleSkills)
	route.Put("/:peopleskillsid", h.UpdateSkill)
	route.Delete("/:peopleskillsid", h.DeleteSkills)
}
