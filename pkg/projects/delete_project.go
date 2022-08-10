package projects

import (
	"strings"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteProject(c *fiber.Ctx) error {
	id := c.Params("projectid")

	var project models.Project

	trim_id := strings.TrimLeft(id, "projectid=")

	if result := h.DB.First(&project, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Where("projectid = ?", trim_id).Delete(&project)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success Deleted!", "data": &project})
}
