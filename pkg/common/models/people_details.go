package models

type PeopleDetails struct {
	PeopleDetailsId    		int  `gorm:"column:peopledetailsid" json:"people_details_id"`
	PeopleId 							int  `gorm:"column:peopleid" json:"people_id"`
	PeopleDetailsDescId 	int  `gorm:"column:peopledetailsdescid" json:"people_details_desc_id"`
	ActiveFlag            bool `gorm:"column:activeflag" json:"is_active"`
}

func (PeopleDetails) TableName() string {
	return "peopledetails"
}

type PeopleDetailsDesc struct {
	PeopleDetailsDescId 	int  `gorm:"column:peopledetailsdescid" json:"people_details_desc_id"`
	PeopleDetailsDesc			string `gorm:"column:peopledetailsdesc" json:"people_details_desc"`
	IsActive							bool `gorm:"column:isactive" json:"is_active"`
}

func (PeopleDetailsDesc) TableName() string {
	return "peopledetailsdesc"
}

type InsertPeopleDetail struct {
	PeopleId 							int  `gorm:"column:peopleid"`
	PeopleDetailsDescId 	int  `gorm:"column:peopledetailsdescid"`
	ActiveFlag            bool `gorm:"column:activeflag"`
}

func (InsertPeopleDetail) TableName() string {
	return "peopledetails"
}