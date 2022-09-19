package models

type Project struct {
	ProjectId   int    `gorm:"primaryKey;column:projectid" json:"id"`
	ProjectName string `gorm:"column:projectdesc" json:"project"`
	ProjectCode string `gorm:"column:projectcode" json:"project_code"`
	IsActive    bool   `gorm:"column:isactive" json:"is_active"`
}

func (Project) TableName() string {
	return "project"
}
