package community

import (
	"strconv"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	requests "github.com/VncntDzn/community-tracker-api/pkg/community/requests"
	communityValidation "github.com/VncntDzn/community-tracker-api/pkg/validations/community"
	"github.com/gofiber/fiber/v2"
)

func (h handler) UpdateCommunity(c *fiber.Ctx) error {

	communityId := c.Params("communityId")

	communityRouteId, conversionError := strconv.Atoi(communityId)
	if conversionError != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "Request not found"})
	}

	updateCommunityRequest := new(requests.CreateCommunityRequest) //reuse create request cause same payload

	c.BodyParser(updateCommunityRequest)

	vadlidationError := communityValidation.ValidateCreateCommunity(*updateCommunityRequest)
	if vadlidationError != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "error": vadlidationError})
	}

	var updatedCommunity models.CreateCommunity
	updatedCommunity.CommunityID = communityRouteId

	updateModel := h.DB.Model(&updatedCommunity).Updates(&models.CreateCommunity{
		CommunityName:    updateCommunityRequest.CommunityName,
		CommunityManager: updateCommunityRequest.CommunityManager,
		CommunityDesc:    updateCommunityRequest.CommunityDesc,
	})
	if updateModel.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": updateModel.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "data": &updatedCommunity, "message": "Community Updated"})
}
