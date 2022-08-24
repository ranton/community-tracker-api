package admin

type CreateAdminRequest struct {
	CognizantId string `validate:"required,number" json:"cognizant_id"`
	Email       string `validate:"required,email" json:"email"`
	AdminName   string `validate:"required,alpha" json:"name"`
}
