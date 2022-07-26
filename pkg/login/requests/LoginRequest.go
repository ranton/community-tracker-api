package login

type LoginRequest struct {
	CognizantId string `json:"cognizant_id"`
	Password    string `json:"password"`
}
