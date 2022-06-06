package models

type SoftDeleteMember struct {
	PeopleID    string `gorm:"column:peopleid" json:"people_id"`
	CognizantID int    `gorm:"column:cognizantid" json:"cognizant_id"`
	LastName    string `gorm:"column:lastname" json:"last_name"`
	FirstName   string `gorm:"column:firstname" json:"first_name"`
	MiddleName  string `gorm:"column:middlename" json:"middle_name"`
	HiredDate   string `gorm:"column:hireddate" json:"hired_date"`
	CSVEmail    string `gorm:"column:csvemail" json:"csv_email"`
	JobLevel    int    `gorm:"column:joblevelid" json:"job_level_id"`
	CommunityID string `gorm:"column:communityid" json:"community_id"`
	IsActive    string `gorm:"column:isactive" json:"is_active"`
}

func (SoftDeleteMember) TableName() string {
	return "communities"
}
