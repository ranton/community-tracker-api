package people

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetWorkState(c *fiber.Ctx) error {
	var work_state []models.WorkState

	if result := h.DB.Find(&work_state); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Success!", "data": &work_state})
}
