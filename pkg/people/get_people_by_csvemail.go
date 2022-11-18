package people

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetPeopleByCSVEmail(c *fiber.Ctx) error {
	var People models.People

	csvemail := c.Params("csv_email")

	if csvemail == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "Not Found"})
	}

	if peopleResult := h.DB.Preload("Community", "isactive = ?", true).Preload("Community.Manager").First(&People, "csvemail = ?", csvemail); peopleResult.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "Not Found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success!", "data": &People})
}
