package peopleDetails

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h handler) AddPeopleDetails(c *fiber.Ctx) error {
	body := models.AddPeopleDetails{
		People: 0,
		PeopleDetailsDesc: 1,
		IsActive: false,
		Gop: false,               
		ExpSettMtg: false,                        
		SignedExpSettDoc: false,               
		MonthlyTPIntroEmail: false,             
		PerfEvalRecurringMtg: false,             
		CommunityComms: false,                
		GopInstEmail: false,                     
		SecondMonthTPMtg: false,              
		ThirdMonthPerfEvalFromProjLead: false,  
		ThirdMonthTPPerfEvalFromProjLead: false,   
		ThirdMonthTPMtg: false,             
		FourthMonthTPMtg: false,              
		FifthTPMonthPerfEvalReqFromProjLead: false,
		FifthMonthPerfEvalFromProjLead: false,
		FifthMonthTPMtg: false,
	}

	// parse body, attach to AddPeopleRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var peopleDetails = models.AddPeopleDetails(body)

	// create transaction for insert
	transactionErr := h.DB.Transaction(func(tx *gorm.DB) error {
		//insert peopledetails
		if createPeopleErr := tx.Create(&peopleDetails).Error; createPeopleErr != nil {
			return createPeopleErr
		}
		return nil
	})

	if transactionErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": transactionErr.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Success! Added Data!", "data": &peopleDetails})
}
