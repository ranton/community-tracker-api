package community_managers

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetManagers(c *fiber.Ctx) error {
	var people []models.People

	if result := h.DB.Table("community").Select("*").Joins("inner join people on people.communityid = community.communityid").Scan(&people); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Success!", "data": &people})
}
