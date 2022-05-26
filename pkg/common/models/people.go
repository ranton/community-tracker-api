package models

type People struct {
	Peopleid    int    `gorm:"column:peopleid" json:"people_id"`
	Cognizantid int    `gorm:"column:cognizantid" json:"cognizantid_id"`
	Lastname    string `gorm:"column:lastname" json:"last_name"`
	Firstname   string `gorm:"column:firstname" json:"first_name"`
	Middlename  string `gorm:"column:middlename" json:"middle_name"`
	Fullname    string `gorm:"column:fullname" json:"full_name"`
	Csvemail    string `gorm:"column:csvemail" json:"csv_email"`
	Hireddate   string `gorm:"column:hireddate" json:"hired_date"`
	Workstateid int    `gorm:"column:workstateid" json:"workstate_id"`
	Joblevelid  int    `gorm:"column:joblevelid" json:"joblevel_id"`
	Projectid   int    `gorm:"column:projectid" json:"project_id"`
	Isactive    bool   `gorm:"column:isactive" json:"is_active"`
}

func (People) TableName() string {
	return "people"
}

type AddPeople struct {
	PeopleID    string `gorm:"column:peopleid" json:"people_id"`
	CognizantID int    `gorm:"column:cognizantid" json:"cognizant_id"`
	LastName    string `gorm:"column:lastname" json:"last_name"`
	FirstName   string `gorm:"column:firstname" json:"first_name"`
	MiddleName  string `gorm:"column:middlename" json:"middle_name"`
	HiredDate   string `gorm:"column:hireddate" json:"hired_date"`
	CSVEmail    string `gorm:"column:csvemail" json:"csv_email"`
	JobLevel    int    `gorm:"column:joblevelid" json:"job_level_id"`
	CommunityID string `gorm:"column:communityid" json:"community_id"`
	IsActive    bool   `gorm:"column:isactive" json:"is_active"`
}

func (AddPeople) TableName() string {
	return "people"
}

type UpdatePeople struct {
	PeopleID    string `gorm:"column:peopleid" json:"people_id"`
	CognizantID int    `gorm:"column:cognizantid" json:"cognizant_id"`
	LastName    string `gorm:"column:lastname" json:"last_name"`
	FirstName   string `gorm:"column:firstname" json:"first_name"`
	MiddleName  string `gorm:"column:middlename" json:"middle_name"`
	HiredDate   string `gorm:"column:hireddate" json:"hired_date"`
	CSVEmail    string `gorm:"column:csvemail" json:"csv_email"`
	JobLevel    int    `gorm:"column:joblevelid" json:"job_level_id"`
	CommunityID string `gorm:"column:communityid" json:"community_id"`
	IsActive    bool   `gorm:"column:isactive" json:"is_active"`
}

func (UpdatePeople) TableName() string {
	return "people"
}
