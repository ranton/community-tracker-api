package community

import (
	"strings"
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

type UpdateCommunityRequestBody struct {
	//CommunityID      int    `gorm:"primaryKey;column:communityid" json:"community_id"`
	CommunityName    string `validate:"required" gorm:"column:communityname" json:"community_name"`
	CommunityManager int    `validate:"required" gorm:"column:communitymgrid" json:"community_manager"`
	CommunityDesc    string `validate:"required" gorm:"column:communitydesc" json:"community_description"`
	Icon             string `gorm:"column:communityicon" json:"icon"`
	IsActive       	 bool   `gorm:"column:isactive" json:"is_active"`
}

func (h handler) UpdateCommunity(c *fiber.Ctx) error {
	id := c.Params("communityid")
	body := UpdateCommunityRequestBody{
		CommunityName:    "",
		CommunityManager: 0,
		CommunityDesc:    "",
		Icon:             "",
		IsActive:					false,
	}

	trim_id := strings.TrimLeft(id, "communityid=")

	// parse body, attach to UpdateCityRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var validate = validator.New()
	validateErr := validate.Struct(body)
	if validateErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "message": validateErr})
	}

	var community models.UpdateCommunity

	community.CommunityName = body.CommunityName
	community.CommunityManager = body.CommunityManager
	community.CommunityDesc = body.CommunityDesc
	community.Icon = body.Icon
	community.IsActive = body.IsActive

	if result := h.DB.First(&community, id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

	} else {

		community.CommunityName = body.CommunityName
		community.CommunityManager = body.CommunityManager
		community.CommunityDesc = body.CommunityDesc
		community.Icon = body.Icon
		community.IsActive = body.IsActive

		mp := make(map[string]interface{})
		mp["communityname"] = body.CommunityName
		mp["communitymgrid"] = body.CommunityManager
		mp["communitydesc"] = body.CommunityDesc
		mp["communityicon"] = body.Icon
		mp["isactive"] = body.IsActive

		h.DB.Model(community).Where("communityid = ?", trim_id).Updates(mp)

		h.DB.Save(&community)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Updated data!", "data": &community})
	}

}
