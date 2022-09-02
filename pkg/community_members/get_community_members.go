package community_members

import (
	"errors"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h handler) GetCommunityMembers(c *fiber.Ctx) error {
	var community_data models.CommunityMembers
	communityId := c.Params("communityId")

	result := h.DB.Where(&models.CommunityMembers{CommunityID: communityId}).Preload("Members", "isactive = ?", true).Preload("Manager").First(&community_data)

	// show 404 error if no community found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "Community does not exist."})
	}

	// check for other errors
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": "Server error."})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success!", "data": &community_data})
}
