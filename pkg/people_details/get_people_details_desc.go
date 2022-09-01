package people_details

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetPeopleDetailsDesc(c *fiber.Ctx) error {
	var peopleDetails []models.PeopleDetailsDesc

	if result := h.DB.Find(&peopleDetails); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success!", "data": &peopleDetails})
}
