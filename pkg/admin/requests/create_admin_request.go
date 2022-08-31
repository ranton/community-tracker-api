package admin

type CreateAdminRequest struct {
	CognizantId string `validate:"required" json:"cognizant_id"`
	Email       string `validate:"required,email" json:"email"`
	AdminName   string `validate:"required" json:"name"`
}
