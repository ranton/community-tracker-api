package models

type Community struct {
	CommunityID              string `gorm:"column:communityid" json:"community_id"`
	CommunityName            string `gorm:"column:communityname" json:"community_name"`
	CommunityManagerPeopleID string `gorm:"column:communitymgrid" json:"community_manager_people_id"`
	Image                    string `gorm:"column:image" json:"image"`
	ColorTheme               string `gorm:"column:colortheme" json:"color_theme"`
	IsActive                 bool   `gorm:"column:isactive" json:"is_active"`
	Manager                  People `gorm:"foreignKey:PeopleID;references:CommunityManagerPeopleID" json:"manager_info"`
}

func (Community) TableName() string {
	return "community"
}

type CommunityMembers struct {
	CommunityID   string   `gorm:"column:communityid" json:"community_id"`
	CommunityName string   `gorm:"column:communityname" json:"community_name"`
	Members       []People `gorm:"foreignKey:CommunityID;references:CommunityID" json:"members"`
}

func (CommunityMembers) TableName() string {
	return "community"
}

type CreateCommunity struct {
	CommunityID      int    `gorm:"primaryKey;column:communityid" json:"community_id"`
	CommunityName    string `gorm:"column:communityname" json:"community_name"`
	CommunityManager int    `gorm:"column:communitymgrid" json:"community_manager"`
	CommunityDesc    string `gorm:"column:communitydesc" json:"community_description"`
}

func (CreateCommunity) TableName() string {
	return "community"
}
