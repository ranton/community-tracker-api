package models

type Cities struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type AddCity struct {
	Name string `json:"name"`
}
