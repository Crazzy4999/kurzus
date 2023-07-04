package models

type Item struct {
	ID         int    `json:"id"`
	MenuId     int64  `json:"menu"`
	Ingerdient string `json:"ingredient"`
}
