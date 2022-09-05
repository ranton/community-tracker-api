package models

type AdminManager struct {
	ID          int    `gorm:"column:communityadminandmanagerid;primaryKey" json:"id"`
	CognizantID string `gorm:"column:cognizantid" json:"cognizant_id"`
	AdminName   string `gorm:"column:communityadminandmanagername" json:"name"`
	Email       string `gorm:"column:csvemail" json:"email"`
	Password    string `gorm:"column:password" json:"-"`
	RoleType    string `gorm:"column:roletype" json:"-"`
	IsActive    bool   `gorm:"column:isactive" json:"active"`
}

func (AdminManager) TableName() string {
	return "communityadminandmanager"
}

type UpdateAdminManager struct {
	AdminName string `gorm:"column:communityadminandmanagername" json:"name"`
	Email     string `gorm:"column:csvemail" json:"email"`
	RoleType  string `gorm:"column:roletype" json:"-"`
	IsActive  bool   `gorm:"column:isactive" json:"active"`
}

func (UpdateAdminManager) TableName() string {
	return "communityadminandmanager"
}


type UpdatePassword struct {
	Password string `gorm:"column:password" json:"password"`
}

func (UpdatePassword) TableName() string {
	return "communityadminandmanager"
}
