package community

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetJobLevel(c *fiber.Ctx) error {
	var job_level []models.JobLevel

	if result := h.DB.Find(&job_level); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Success!", "data": &job_level})
}
