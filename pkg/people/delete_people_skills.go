package people

import (
	"fmt"
	"strings"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type DeletePeopleRequestBody struct {
	Isactive bool `gorm:"column:isactive" json:"is_active"`
}

func (h handler) DeletePeople(c *fiber.Ctx) error {
	id := c.Params("peopleskillsid")
	body := UpdatePeopleRequestBody{
		Isactive: false,
	}

	trim_id := strings.TrimLeft(id, "peopleskillsid=")

	// parse body, attach to UpdateCityRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var peopleskills models.Deleteskills

	peopleskills.IsActive = body.Isactive

	if result := h.DB.First(&peopleskills, id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

	} else {

		peopleskills.IsActive = body.Isactive

		h.DB.Model(&peopleskills).Where("peopleskillsid = ?", trim_id).Update("isactive", body.Isactive)

		h.DB.Save(&peopleskills)
		fmt.Println(result)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Updated data!", "data": &peopleskills})
	}

}
