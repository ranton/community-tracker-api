package cities

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type DeleteCityRequestBody struct {
	Name string `json:"name"`
}

func (h handler) DeleteCity(c *fiber.Ctx) error {
	id := c.Params("id")

	var city models.Cities

	if result := h.DB.First(&city, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// delete city from db
	h.DB.Delete(&city)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Success Deleted!", "data": &city})
}
