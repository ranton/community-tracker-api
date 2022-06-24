package community_members

import (
	"strings"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type SoftDeleteRequestBody struct {
	isActive string `json:"is_active"`
}

func (h handler) SoftDeleteMember(c *fiber.Ctx) error {
	id := c.Params("people_id")
	body := SoftDeleteRequestBody{
		isActive: "0",
	}

	trim_id := strings.TrimLeft(id, "people_id=")

	// parse body, attach to SoftDeleteRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var member models.SoftDeleteMember

	member.IsActive = body.isActive

	if result := h.DB.First(&member, id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	} else {

		member.IsActive = body.isActive

		h.DB.Model(&member).Where("people_id = ?", trim_id).Update("isActive", body.isActive)

		h.DB.Save(&member)
		//fmt.Println(result)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Updated data!", "data": &member})
	}

}
