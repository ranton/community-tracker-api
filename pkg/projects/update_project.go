package projects

import (
	"strings"

	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

type UpdateProjectDetails struct {
	/* ProjectId     int    `gorm:"primaryKey;column:project_id" json:"project_id"` */
	ProjectName string `validate:"required" gorm:"column:projectdesc" json:"project_name"`
	ProjectCode string `validate:"required" gorm:"column:projectcode" json:"project_code"`
	IsActive    bool   `gorm:"column:isactive" json:"is_active"`
}

func (h handler) UpdateProject(c *fiber.Ctx) error {
	id := c.Params("projectid")
	body := UpdateProjectDetails{
		ProjectName: "",
		ProjectCode: "",
		IsActive:    false,
	}

	trim_id := strings.TrimLeft(id, "projectid=")

	// parse body, attach to UpdateCityRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var validate = validator.New()
	validateErr := validate.Struct(body)
	if validateErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"status": fiber.StatusUnprocessableEntity, "message": validateErr})
	}

	var project models.Project

	project.ProjectName = body.ProjectName
	project.ProjectCode = body.ProjectCode
	project.IsActive = body.IsActive

	if result := h.DB.First(&project, id); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

	} else {

		project.ProjectName = body.ProjectName
		project.ProjectCode = body.ProjectCode
		project.IsActive = body.IsActive

		mp := make(map[string]interface{})
		mp["projectdesc"] = body.ProjectName
		mp["projectcode"] = body.ProjectCode
		mp["is_active"] = body.IsActive

		h.DB.Model(project).Where("projectid = ?", trim_id).Updates(mp)

		h.DB.Save(&project)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Updated data!", "data": &project})
	}

}
