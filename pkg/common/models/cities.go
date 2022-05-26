package models

type Cities struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
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
