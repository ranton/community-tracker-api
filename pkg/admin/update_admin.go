package admin

import (
	admin "github.com/VncntDzn/community-tracker-api/pkg/admin/requests"
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

func (h handler) UpdateAdminDetails(c *fiber.Ctx) error {
	id := c.Params("communityadminandmanagerid")
	body := admin.UpdateAdminRequest{
		CommunityManagerAdminName: "",
		CSV_EMAIL:                 "",
		ROLE_TYPE:                 "",
		IS_ACTIVE:                 true,
	}

	// parse body, attach to UpdateCityRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var validate = validator.New()
	validateErr := validate.Struct(body)
	if validateErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "message": validateErr})
	}

	var community models.UpdateAdminManager

	community.AdminName = body.CommunityManagerAdminName
	community.Email = body.CSV_EMAIL
	community.RoleType = body.ROLE_TYPE
	community.IsActive = body.IS_ACTIVE
	if result := h.DB.First(&community, "communityadminandmanagerid = ?", id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

	} else {

		community.AdminName = body.CommunityManagerAdminName
		community.Email = body.CSV_EMAIL
		community.RoleType = body.ROLE_TYPE
		community.IsActive = body.IS_ACTIVE
		mp := make(map[string]interface{})
		mp["communityadminandmanagername"] = body.CommunityManagerAdminName
		mp["csvemail"] = body.CSV_EMAIL
		mp["roletype"] = body.ROLE_TYPE
		mp["isactive"] = body.IS_ACTIVE

		h.DB.Model(community).Where("communityadminandmanagerid = ?", id).Updates(mp)

		h.DB.Save(&community)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Updated data!", "data": &community})
	}

}
