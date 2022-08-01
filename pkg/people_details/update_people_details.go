package peopleDetails

import (
	"github.com/VncntDzn/community-tracker-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) UpdatePeopleDetails(c *fiber.Ctx) error {
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
	if result := h.DB.Where("peopleid = ?", peopleDetails.People).First(&models.AddPeopleDetails{}); result.Error != nil {

		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

	} else {

		h.DB.Model(&models.AddPeopleDetails{}).Where("peopleid = ?", peopleDetails.People).Updates(&peopleDetails)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Updated data!", "data": &peopleDetails})
	}
}
