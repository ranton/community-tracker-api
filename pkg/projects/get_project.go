package projects

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetProject(c *fiber.Ctx) error {
	id := c.Params("projectid")
	var project models.Project

	if result := h.DB.First(&project, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success!", "data": &project})
}
