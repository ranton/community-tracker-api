package projects

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	requests "github.com/VncntDzn/community-tracker-api/pkg/projects/requests"
	projectValidation "github.com/VncntDzn/community-tracker-api/pkg/validations/projects"
	"github.com/gofiber/fiber/v2"
)

func (h handler) CreateProject(c *fiber.Ctx) error {
	projectRequest := new(requests.CreateProjectRequest)

	c.BodyParser(projectRequest)

	validateErr := projectValidation.ValidateCreateCommunity(*projectRequest)
	if validateErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "error": validateErr})
	}

	//@Todo add validation for project name

	projectData := models.Project{
		ProjectName: projectRequest.ProjectName,
		ProjectLead: projectRequest.ProjectLead,
		IsActive:    true,
	}

	projectResult := h.DB.Create(&projectData)
	if projectResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "error": projectResult.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": fiber.StatusOK, "message": "Success!", "data": &projectRequest})
}
