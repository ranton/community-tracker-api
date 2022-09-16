package people

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (h handler) GetPeople(c *fiber.Ctx) error {
	var people []models.People

	var searchCriteria = strings.ToLower(strings.TrimSpace(c.Query("searchCriteria")))

	if searchCriteria == "" {
		if result := h.DB.Order("lower(fullname)").Where("isactive = ?", true).Find(&people); result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
		}
	} else {
		if result := h.DB.Order("lower(fullname)").Where(h.DB.Where("isactive = ?", true).Where(h.DB.Where("lower(fullname) like ?", "%" + searchCriteria + "%").Or("cast(cognizantid as text) like ?", "%" + searchCriteria + "%"))).Find(&people); result.Error != nil {
			return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success!", "data": &people})
}
