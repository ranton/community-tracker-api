package member_skills

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

	routes := app.Group("/api/members")
	routes.Post("/:memberId/skills", h.AddMemberSkill)
}
