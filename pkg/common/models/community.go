package models

type Community struct {
	CommunityID              string `gorm:"column:communityid" json:"community_id"`
	CommunityName            string `gorm:"column:communityname" json:"community_name"`
	CommunityManagerPeopleID string `gorm:"column:communitymgrid" json:"community_manager_people_id"`
	Image                    string `gorm:"column:image" json:"image"`
	ColorTheme               string `gorm:"column:colortheme" json:"color_theme"`
	IsActive                 bool   `gorm:"column:isactive" json:"is_active"`
}

func (Community) TableName() string {
	return "community"
}
