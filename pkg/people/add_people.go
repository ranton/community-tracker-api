package people

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddPeopleRequestBody struct {
	Cognizantid    int    `gorm:"column:cognizantid" json:"cognizantid_id"`
	Lastname       string `gorm:"column:lastname" json:"last_name"`
	Firstname      string `gorm:"column:firstname" json:"first_name"`
	Middlename     string `gorm:"column:middlename" json:"middle_name"`
	Fullname       string `gorm:"column:fullname" json:"full_name"`
	Csvemail       string `gorm:"column:csvemail" json:"csv_email"`
	Hireddate      string `gorm:"column:hireddate" json:"hired_date"`
	Communityid    int    `gorm:"column:communityid" json:"community_id"`
	Workstateid    int    `gorm:"column:workstateid" json:"workstate_id"`
	Joblevelid     int    `gorm:"column:joblevelid" json:"joblevel_id"`
	Projectid      int    `gorm:"column:projectid" json:"project_id"`
	Isactive       bool   `gorm:"column:isactive" json:"is_active"`
	Isprobationary bool   `gorm:"column:isprobationary" json:"is_probationary"`
}

func (h handler) AddPeople(c *fiber.Ctx) error {
	body := AddPeopleRequestBody{
		Cognizantid:    0,
		Lastname:       "",
		Firstname:      "",
		Middlename:     "",
		Fullname:       "",
		Csvemail:       "",
		Hireddate:      "",
		Communityid:    0,
		Workstateid:    0,
		Joblevelid:     0,
		Projectid:      0,
		Isactive:       true,
		Isprobationary: false,
	}

	// parse body, attach to AddPeopleRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var people models.Add_People

	people.Cognizantid = body.Cognizantid
	people.Lastname = body.Lastname
	people.Firstname = body.Firstname
	people.Middlename = body.Middlename
	people.Fullname = body.Fullname
	people.Hireddate = body.Hireddate
	people.Csvemail = body.Csvemail
	people.Workstateid = body.Workstateid
	people.Communityid = body.Communityid
	people.Joblevelid = body.Joblevelid
	people.Projectid = body.Projectid
	people.Isactive = body.Isactive
	people.Isprobationary = body.Isprobationary

	// insert new db entry
	if result := h.DB.Create(&people); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Success! Added Data!", "data": &people})
}
