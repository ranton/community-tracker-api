package models

type Community struct {
	CommunityID              string `json:"community_id"`
	CommunityName            string `json:"community_name"`
	CommunityManagerPeopleID string `json:"community_manager_people_id"`
	Image                    string `json:"image"`
	ColorTheme               string `json:"color_theme"`
	IsActive                 bool   `json:"is_active"`
}
