package community

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

type AddCommunityBody struct {
	CommunityID      int    `gorm:"primaryKey;column:communityid" json:"community_id"`
	CommunityName    string `validate:"required" gorm:"column:communityname" json:"community_name"`
	CommunityManager int    `gorm:"column:communitymgrid" json:"community_manager"`
	CommunityDesc    string `validate:"required" gorm:"column:communitydesc" json:"community_description"`
	Icon             string `gorm:"column:communityicon" json:"icon"`
	IsActive		 bool 	`gorm:"column:isactive" json:"is_active"`
}

func (h handler) AddCommunity(c *fiber.Ctx) error {

	body := AddCommunityBody{
		CommunityName:    "",
		CommunityManager: 0,
		CommunityDesc:    "",
		Icon:             "",
		IsActive:		  true,
	}

	// parse body, attach to struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var validate = validator.New()
	validateErr := validate.Struct(body)
	if validateErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "message": validateErr})
	}

	var community models.CreateCommunity

	community.CommunityName = body.CommunityName
	community.CommunityManager = body.CommunityManager
	community.CommunityDesc = body.CommunityDesc
	community.Icon = body.Icon
	community.IsActive = body.IsActive

	// insert new records
	if result := h.DB.Create(&community); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Added New Data!", "data": &community})
}
