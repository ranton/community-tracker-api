package login

type LoginRequest struct {
	CognizantId string `validate:"required" json:"cognizant_id"`
	Password    string `validate:"required" json:"password"`
}
