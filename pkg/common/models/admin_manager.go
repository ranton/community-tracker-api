package models

type AdminManager struct {
	ID          uint   `gorm:"column:communityadminandmanagerid;primaryKey" json:"id"`
	CognizantID string `gorm:"column:cognizantid" json:"cognizant_id"`
	AdminName   string `gorm:"column:communityadminandmanagername" json:"name"`
	Email       string `gorm:"column:csvemail" json:"email"`
	Password    string `gorm:"column:password" json:"password"`
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
