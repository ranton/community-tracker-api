package cities

import (
	"strings"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateCityRequestBody struct {
	Name string `json:"name"`
	//City_id int    `json:"city_id"`
}

func (h handler) UpdateCity(c *fiber.Ctx) error {
	id := c.Params("city_id")
	body := UpdateCityRequestBody{
		Name: "",
	}

	trim_id := strings.TrimLeft(id, "city_id=")

	// parse body, attach to UpdateCityRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var city models.UpdateCity

	city.Name = body.Name

	if result := h.DB.First(&city, id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	} else {

		city.Name = body.Name

		h.DB.Model(&city).Where("city_id = ?", trim_id).Update("name", body.Name)

		h.DB.Save(&city)
		//fmt.Println(result)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Updated data!", "data": &city})
	}

}
