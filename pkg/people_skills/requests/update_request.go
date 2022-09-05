package people_skills

type UpdateSkillRequest struct {
	Peopleskillsdesc string `validate:"required" gorm:"column:peopleskillsdesc" json:"description"`
	IsActive         bool   `gorm:"column:isactive" json:"is_active"`
}
