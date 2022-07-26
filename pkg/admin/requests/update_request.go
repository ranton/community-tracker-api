package admin

type UpdateAdminRequest struct {
	CommunityManagerAdminName string `gorm:"column:communityadminandmanagername" json:"name"`
	CSV_EMAIL                 string `gorm:"column:csvemail" json:"csv_email"`
	ROLE_TYPE                 string `gorm:"column:roletype" json:"role_type"`
	IS_ACTIVE                 bool   `gorm:"column:isactive" json:"is_active"`
}
