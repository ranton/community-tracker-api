package models

type Cities struct {
	Name string `json:"name"`
	//City_id int    `json:"city_id" primaryKey;autoIncrement:true`
	City_id int `gorm:"column:city_id" json:"city_id"`
}

type AddCity struct {
	Name string `json:"name"`
}

func (AddCity) TableName() string {
	return "cities"
}

type UpdateCity struct {
	Name string `json:"name"`
}

func (UpdateCity) TableName() string {
	return "cities"
}
