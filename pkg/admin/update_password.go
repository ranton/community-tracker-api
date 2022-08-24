package admin

import (
	admin "github.com/VncntDzn/community-tracker-api/pkg/admin/requests"
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/VncntDzn/community-tracker-api/pkg/utils/hash"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (h handler) UpdatePassword(c *fiber.Ctx) error {
	id := c.Params("communityadminandmanagerid")
	body := admin.UpdatePasswordRequest{
		Password: "",
		NewPassword: "",
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var validate = validator.New()
	validateErr := validate.Struct(body)
	if validateErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "message": validateErr})
	}

	var community models.AdminManager
	var passwordModel models.UpdatePassword
	if result := h.DB.First(&community, "communityadminandmanagerid = ?", id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

	} else {
		if !hash.Check(body.Password, community.Password) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "Invalid Password"})
		}

		if hash.Check(body.NewPassword, community.Password) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "Same Password"})
		}

		mp := make(map[string]interface{})
		mp["password"] = hash.Make(body.NewPassword)

		h.DB.Model(passwordModel).Where("communityadminandmanagerid = ?", id).Updates(mp)

		h.DB.Save(&passwordModel)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Updated data!"})
	}

}
