package admin

type UpdatePasswordRequest struct {
	NewPassword string `gorm:"column:password" json:"newpassword"`
}
