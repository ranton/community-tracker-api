package community

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	requests "github.com/VncntDzn/community-tracker-api/pkg/community/requests"
	communityValidation "github.com/VncntDzn/community-tracker-api/pkg/validations/community"
	"github.com/gofiber/fiber/v2"
)

func (h handler) CreateCommunity(c *fiber.Ctx) error {

	communityRequest := new(requests.CreateCommunityRequest)

	c.BodyParser(communityRequest)

	err := communityValidation.ValidateCreateCommunity(*communityRequest)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "error": err})
	}

	community := models.CreateCommunity{
		CommunityName:    communityRequest.CommunityName,
		CommunityManager: communityRequest.CommunityManager,
		CommunityDesc:    communityRequest.CommunityDesc,
	}

	communityModel := h.DB.Create(&community)
	if communityModel.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": communityModel.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "data": &community, "message": "Community Created"})
}
