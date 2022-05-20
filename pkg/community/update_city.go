package community

import (
	"fmt"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateCityRequestBody struct {
	Name string `json:"name"`
}

func (h handler) UpdateCity(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateCityRequestBody{}

	// parse body, attach to UpdateCityRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var city models.Cities

	city.Name = body.Name

	if result := h.DB.First(&city, id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	} else {
		city.Name = body.Name
		h.DB.Save(&city)
		fmt.Println(result)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Updated data!", "data": &city})
	}

}
