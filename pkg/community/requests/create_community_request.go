package community

type CreateCommunityRequest struct {
	CommunityName    string `validate:"required" json:"community_name"`
	CommunityManager int    `json:"community_manager"`
	CommunityDesc    string `validate:"required" json:"community_description"`
}
