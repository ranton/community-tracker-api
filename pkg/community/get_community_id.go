package community

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetCommunityById(c *fiber.Ctx) error {
	var Community models.Community

	id := c.Params("communityid")

	h.DB.First(&Community, "communityid = ?", id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": &Community})
}
