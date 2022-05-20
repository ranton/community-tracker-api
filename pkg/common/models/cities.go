package models

type Cities struct {
	Name string `json:"name"`
	ID   int    `json:"id" primaryKey;autoIncrement:true`
}

type AddCity struct {
	Name string `json:"name"`
}
