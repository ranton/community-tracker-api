package people_skills

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetPeopleSkills(c *fiber.Ctx) error {
	var peopleskills []models.Peopleskills

	if result := h.DB.Find(&peopleskills); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Success!", "data": &peopleskills})
}
