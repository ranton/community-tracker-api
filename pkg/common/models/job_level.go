package models

type JobLevel struct {
	JobLevelID          int    `gorm:"column:joblevelid" json:"job_level_id"`
	JobLevelDescription string `gorm:"column:jobleveldesc" json:"job_level_desc"`
	IsActive            bool   `gorm:"column:isactive" json:"is_active"`
}

func (JobLevel) TableName() string {
	return "joblevel"
}