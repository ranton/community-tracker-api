package admin

type UpdatePasswordRequest struct {
	Password string `validate:"required" gorm:"column:password" json:"password"`
	NewPassword string `validate:"required" gorm:"column:password" json:"newpassword"`
}
