package member_skills

type SkillsListRequest struct {
	Skills string `validate:"required" json:"skills"`
}
