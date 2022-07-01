package people

import (
	"fmt"
	"strings"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdatePeopleSkillsRequestBody struct {
	Peopleskillsdesc string `gorm:"column:peopleskillsdesc" json:"peopleskills_desc"`
	IsActive         bool   `gorm:"column:isactive" json:"is_active"`
}

func (h handler) UpdatePeopleSkills(c *fiber.Ctx) error {
	id := c.Params("peopleskillsid")
	body := UpdatePeopleSkillsRequestBody{
		Peopleskillsdesc: "",
		IsActive:         true,
	}

	trim_id := strings.TrimLeft(id, "peopleskillsid=")

	// parse body, attach to UpdateCityRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var peopleskills models.Update_People_Skills

	peopleskills.Peopleskillsdesc = body.Peopleskillsdesc

	if result := h.DB.First(&peopleskills, id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

	} else {

		peopleskills.Peopleskillsdesc = body.Peopleskillsdesc
		peopleskills.IsActive = body.IsActive

		h.DB.Model(&peopleskills).Where("peopleskillsid = ?", trim_id).Update("peopleskillsdesc", body.Peopleskillsdesc)

		h.DB.Save(&peopleskills)
		fmt.Println(result)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Updated data!", "data": &peopleskills})
	}
}
