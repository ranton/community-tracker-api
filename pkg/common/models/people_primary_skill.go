package models

type PeoplePrimarySkills struct {
	PeopleId    int `gorm:"column:peopleid" json:"-"`
	PeopleSkill int `gorm:"column:peopleskillsid" json:"skill_id"`
}

func (PeoplePrimarySkills) TableName() string {
	return "peopleprimaryskills"
}

type InsertPrimarySkill struct {
	PeopleId    int  `gorm:"column:peopleid" `
	PeopleSkill int  `gorm:"column:peopleskillsid" `
	IsActive    bool `gorm:"column:isactive"`
}

func (InsertPrimarySkill) TableName() string {
	return "peopleprimaryskills"
}
