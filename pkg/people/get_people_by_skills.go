package people
import (
    "strings"
    "strconv"
    "github.com/gofiber/fiber/v2"
)
type GetPeopleRequestBody struct {
    Skills string `json:"skills"`
}
type PeopleWithSkill struct {
    Fullname    string  `gorm:"column:fullname"`
    PeopleSkillsId  int `gorm:"column:peopleskillsid"`
    PeopleSkillsDesc    string  `gorm:"column:peopleskillsdesc"`
	ProjectId    int  `gorm:"column:projectid"`
}
type PeopleWithSkillResponseBody struct {
    Fullname    string  `gorm:"column:fullname" json:"full_name"`
    Skills  []string `gorm:"column:skills" json:"skills"`
	ProjectId int `gorm:"column:projectid" json:"project_id"`
}
func (h handler) GetPeopleBySkills(c *fiber.Ctx) error {
    body := GetPeopleRequestBody{
        Skills: "",
    }
    if err := c.BodyParser(&body); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, err.Error())
    }
    var s []string
    s = strings.Split(body.Skills, ",");
    var t2 = make([]int, len(s))
    for idx, i := range s {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        t2[idx] = j
    }
    var people []PeopleWithSkill
    //subquery between peopleprimaryskills and peopleskills
    sub := h.DB.Table("peopleprimaryskills").Select("peopleprimaryskills.peopleid, peopleskills.peopleskillsid, peopleskills.peopleskillsdesc").Joins("inner join peopleskills on peopleprimaryskills.peopleskillsid = peopleskills.peopleskillsid")
    if result := h.DB.Table("people").Select("projectid, fullname, sub.peopleskillsid, sub.peopleskillsdesc").Joins("inner join (?) as sub on sub.peopleid = people.peopleid", sub).Where("sub.peopleskillsid IN (?)", t2).Scan(&people); result.Error != nil {
        return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
    }
    occurred := map[string]bool{}
	var currentProject int
    var currentMember string
    currentSkills := []string{}
    var members []PeopleWithSkillResponseBody
    for index, p := range people {
        if occurred[p.Fullname] != true {
            occurred[p.Fullname] = true
            //push prev member
            m := PeopleWithSkillResponseBody {
                Fullname: currentMember,
                Skills: currentSkills,
				ProjectId: currentProject,
            }
            //dont allow empty element to be pushed
            if index != 0 {
                members = append(members, m);
            }
            //reset
            currentMember = p.Fullname
            currentSkills = nil
			currentProject = p.ProjectId
        }
        currentSkills = append(currentSkills, p.PeopleSkillsDesc)
        //push last element before ending the loop
        if len(people)-1 == index {
            m := PeopleWithSkillResponseBody {
                Fullname: currentMember,
                Skills: currentSkills,
				ProjectId: currentProject,
            }
            members = append(members, m);
        }
    }
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": fiber.StatusCreated, "message": "Success!", "data": &members})
}