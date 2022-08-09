package people

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
	peopleRoutes := app.Group("/api/people")
	peopleRoutes.Get("/", h.GetPeople)
	peopleRoutes.Get("/workstate", h.GetWorkState)
	peopleRoutes.Get("/:people_id", h.GetPeopleById)
	peopleRoutes.Post("/", h.AddPeople)
	peopleRoutes.Put("/:peopleid", h.UpdatePeople)
	peopleRoutes.Post("/skills", h.GetPeopleBySkills)
}
