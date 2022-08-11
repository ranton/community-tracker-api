package models

type WorkState struct {
	WorkStateId          int    `gorm:"column:workstateid" json:"work_state_id"`
	WorkStateDescription string `gorm:"column:workstatedescription" json:"work_state_desc"`
	IsActive            bool   `gorm:"column:isactive" json:"is_active"`
}

func (WorkState) TableName() string {
	return "workstate"
}