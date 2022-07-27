package admin

import (
	admin "github.com/VncntDzn/community-tracker-api/pkg/admin/requests"
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/VncntDzn/community-tracker-api/pkg/utils/hash"
	"github.com/gofiber/fiber/v2"
)

func (h handler) UpdatePassword(c *fiber.Ctx) error {
	id := c.Params("communityadminandmanagerid")
	body := admin.UpdatePasswordRequest{
		NewPassword: "",
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var community models.AdminManager
	var passwordModel models.UpdatePassword
	if result := h.DB.First(&community, "communityadminandmanagerid = ?", id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

	} else {
		if hash.Check(body.NewPassword, community.Password) {
			return fiber.NewError(fiber.StatusBadRequest, "Same password")
		}

		mp := make(map[string]interface{})
		mp["password"] = hash.Make(body.NewPassword)

		h.DB.Model(passwordModel).Where("communityadminandmanagerid = ?", id).Updates(mp)

		h.DB.Save(&passwordModel)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Updated data!"})
	}

}
