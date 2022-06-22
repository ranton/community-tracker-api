package member_skills

import (
	"strconv"
	"strings"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	requests "github.com/VncntDzn/community-tracker-api/pkg/member_skills/requests"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h handler) AddMemberSkill(c *fiber.Ctx) error {
	skillsRequest := new(requests.SkillsListRequest)
	c.BodyParser(&skillsRequest)

	skillSet := strings.Split(skillsRequest.Skills, ",")

	memberId := c.Params("memberId")
	parsedMemberId, parseError := strconv.Atoi(memberId)
	if parseError != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": fiber.StatusNotFound, "message": "Resource not found"})
	}
	var batchSkills []*models.InsertPrimarySkill
	for _, skillId := range skillSet {
		parsedSkill, _ := strconv.Atoi(skillId)
		var skillItem models.InsertPrimarySkill
		skillItem.PeopleId = parsedMemberId
		skillItem.PeopleSkill = parsedSkill
		skillItem.IsActive = true
		batchSkills = append(batchSkills, &skillItem)
	}

	transactionErr := h.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&batchSkills).Error; err != nil {
			// return any error will rollback
			return err
		}
		return nil
	})
	if transactionErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": transactionErr.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success!"})
}
