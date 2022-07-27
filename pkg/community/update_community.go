package community

import (
	"strings"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateCommunityRequestBody struct {
	//CommunityID      int    `gorm:"primaryKey;column:communityid" json:"community_id"`
	CommunityName    string `gorm:"column:communityname" json:"community_name"`
	CommunityManager int    `gorm:"column:communitymgrid" json:"community_manager"`
	CommunityDesc    string `gorm:"column:communitydesc" json:"community_description"`
	Icon             string `gorm:"column:communityicon" json:"icon"`
}

func (h handler) UpdateCommunity(c *fiber.Ctx) error {
	id := c.Params("communityid")
	body := UpdateCommunityRequestBody{
		CommunityName:    "",
		CommunityManager: 0,
		CommunityDesc:    "",
		Icon:             "",
	}

	trim_id := strings.TrimLeft(id, "communityid=")

	// parse body, attach to UpdateCityRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var community models.UpdateCommunity

	community.CommunityName = body.CommunityName
	community.CommunityManager = body.CommunityManager
	community.CommunityDesc = body.CommunityDesc
	community.Icon = body.Icon

	if result := h.DB.First(&community, id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

	} else {

		community.CommunityName = body.CommunityName
		community.CommunityManager = body.CommunityManager
		community.CommunityDesc = body.CommunityDesc
		community.Icon = body.Icon

		mp := make(map[string]interface{})
		mp["communityname"] = body.CommunityName
		mp["communitymgrid"] = body.CommunityManager
		mp["communitydesc"] = body.CommunityDesc
		mp["icon"] = body.Icon

		h.DB.Model(community).Where("communityid = ?", trim_id).Updates(mp)

		h.DB.Save(&community)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Updated data!", "data": &community})
	}

}
