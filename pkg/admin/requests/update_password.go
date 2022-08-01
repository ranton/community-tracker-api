package admin

type UpdatePasswordRequest struct {
	Password string `gorm:"column:password" json:"password"`
	NewPassword string `gorm:"column:password" json:"newpassword"`
}
