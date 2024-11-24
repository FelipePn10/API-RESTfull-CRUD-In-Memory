package models

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required"`
	Value int    `json:"value" validate:"required,min=1"`
}
