package models

type Item struct {
	ID         int    `json:"id"`
	MenuId     int    `json:"menu"`
	Ingerdient string `json:"ingredient"`
}
