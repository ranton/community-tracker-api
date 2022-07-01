package models

type Peopleskills struct {
	Peopleskillsid   int    `gorm:"column:peopleskillsid" json:"peopleskills_id"`
	Peopleskillsdesc string `gorm:"column:peopleskillsdesc" json:"peopleskills_desc"`
	IsActive         bool   `gorm:"column:isactive" json:"is_active"`
}

func (Peopleskills) TableName() string {
	return "peopleskills"
}

type Add_People_Skills struct {
	Peopleskillsdesc string `gorm:"column:peopleskillsdesc" json:"peopleskills_desc"`
	IsActive         bool   `gorm:"column:isactive" json:"is_active"`
}

func (Add_People_Skills) TableName() string {
	return "peopleskills"
}

type Deleteskills struct {
	IsActive bool `gorm:"column:isactive" json:"is_active"`
}

func (Deleteskills) TableName() string {
	return "peopleskills"
}

type Update_People_Skills struct {
	Peopleskillsdesc string `gorm:"column:peopleskillsdesc" json:"peopleskills_desc"`
	IsActive         bool   `gorm:"column:isactive" json:"is_active"`
}

func (Update_People_Skills) TableName() string {
	return "peopleskills"
}

type SkillSet struct {
	Peopleskillsid   int    `gorm:"column:peopleskillsid" json:"id"`
	Peopleskillsdesc string `gorm:"column:peopleskillsdesc" json:"description"`
}

func (SkillSet) TableName() string {
	return "peopleskills"
}

type Hard_Delete_Skills struct {
	Peopleskillsid   int    `gorm:"column:peopleskillsid" json:"id"`
	Peopleskillsdesc string `gorm:"column:peopleskillsdesc" json:"description"`
	IsActive         bool   `gorm:"column:isactive" json:"is_active"`
}

func (Hard_Delete_Skills) TableName() string {
	return "peopleskills"
}
