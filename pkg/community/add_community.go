package community

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddCommunityBody struct {
	CommunityID      int    `gorm:"primaryKey;column:communityid" json:"community_id"`
	CommunityName    string `gorm:"column:communityname" json:"community_name"`
	CommunityManager int    `gorm:"column:communitymgrid" json:"community_manager"`
	CommunityDesc    string `gorm:"column:communitydesc" json:"community_description"`
	Icon             string `gorm:"column:icon" json:"icon"`
}

func (h handler) AddCommunity(c *fiber.Ctx) error {

	body := AddCommunityBody{
		CommunityName:    "",
		CommunityManager: 0,
		CommunityDesc:    "",
		Icon:             "",
	}

	// parse body, attach to struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var community models.CreateCommunity

	community.CommunityName = body.CommunityName
	community.CommunityManager = body.CommunityManager
	community.CommunityDesc = body.CommunityDesc
	community.Icon = body.Icon

	// insert new records
	if result := h.DB.Create(&community); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Added New Data!", "data": &community})
}
