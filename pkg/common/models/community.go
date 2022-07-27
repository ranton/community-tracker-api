package models

type Community struct {
	CommunityID    int    `gorm:"column:communityid" json:"community_id"`
	CommunityName  string `gorm:"column:communityname" json:"community_name"`
	CommunityDesc  string `gorm:"column:communitydesc" json:"community_description"`
	CommunityMgrID string `gorm:"column:communitymgrid" json:"community_manager"`
	Icon           string `gorm:"column:icon" json:"icon"`
	//	Manager                  People `gorm:"foreignKey:PeopleID;references:CommunityManagerPeopleID" json:"manager_info"`
}

func (Community) TableName() string {
	return "community"
}

type CommunityMembers struct {
	CommunityID              string   `gorm:"column:communityid" json:"community_id"`
	CommunityName            string   `gorm:"column:communityname" json:"community_name"`
	CommunityManagerPeopleID string   `gorm:"column:communitymgrid" json:"-"`
	Members                  []People `gorm:"foreignKey:Communityid;references:CommunityID" json:"members"`
	Manager                  People   `gorm:"foreignKey:Peopleid;references:CommunityManagerPeopleID" json:"manager"`
}

func (CommunityMembers) TableName() string {
	return "community"
}

type CreateCommunity struct {
	CommunityID      int    `gorm:"primaryKey;column:communityid" json:"community_id"`
	CommunityName    string `gorm:"column:communityname" json:"community_name"`
	CommunityManager int    `gorm:"column:communitymgrid" json:"community_manager"`
	CommunityDesc    string `gorm:"column:communitydesc" json:"community_description"`
	Icon             string `gorm:"column:icon" json:"icon"`
}

func (CreateCommunity) TableName() string {
	return "community"
}

type UpdateCommunity struct {
	CommunityID      int    `gorm:"primaryKey;column:communityid" json:"community_id"`
	CommunityName    string `gorm:"column:communityname" json:"community_name"`
	CommunityManager int    `gorm:"column:communitymgrid" json:"community_manager"`
	CommunityDesc    string `gorm:"column:communitydesc" json:"community_description"`
	Icon             string `gorm:"column:icon" json:"icon"`
}

func (UpdateCommunity) TableName() string {
	return "community"
}

type CommunityWithMembersPercentage struct {
	CommunityID    int    `gorm:"column:communityid" json:"community_id"`
	CommunityName  string `gorm:"column:communityname" json:"community_name"`
	CommunityDesc  string `gorm:"column:communitydesc" json:"community_description"`
	CommunityIcon          string `gorm:"column:communityicon" json:"icon"`
	Percentage int `gorm:"column:percentage" json:"percentage"`
	ManagerFullName string `gorm:"column:fullname" json:"manager_full_name"`
}

func (CommunityWithMembersPercentage) TableName() string {
	return "community"
}

