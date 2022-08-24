package people_skills

import (
	skillRequest "github.com/VncntDzn/community-tracker-api/pkg/people_skills/requests"
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

func (h handler) UpdateSkill(c *fiber.Ctx) error {
	id := c.Params("peopleskillsid")
	body := skillRequest.UpdateSkillRequest{
		Peopleskillsdesc: "",
		IsActive: true,
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var validate = validator.New()
	validateErr := validate.Struct(body)
	if validateErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "message": validateErr})
	}

	var skill models.UpdateSkill

	skill.Peopleskillsdesc = body.Peopleskillsdesc
	skill.IsActive = body.IsActive

	if result := h.DB.First(&skill, "peopleskillsid = ?", id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

	} else {
		mp := make(map[string]interface{})
		mp["peopleskillsdesc"] = body.Peopleskillsdesc
		mp["isactive"] = body.IsActive

		h.DB.Model(skill).Where("peopleskillsid = ?", id).Updates(mp)

		h.DB.Save(&skill)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Updated data!", "data": &skill})
	}

}
