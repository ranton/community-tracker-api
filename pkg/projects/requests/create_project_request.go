package projects

type CreateProjectRequest struct {
	ProjectName 	string `validate:"required" json:"project"`
	ProjectCode 	string `validate:"required" json:"project_code"`
	// ProjectLead int    ` validate:"required" json:"lead"`
	IsActive			bool   `json:"is_active"`
}
