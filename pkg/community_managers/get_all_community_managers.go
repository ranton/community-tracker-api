package community_managers

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetCommunityManagers(c *fiber.Ctx) error {
	var managers []models.AdminManager

	if result := h.DB.Order("lower(communityadminandmanagername)").Where("isactive = ?", true).Find(&managers); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success!", "data": &managers})
}
