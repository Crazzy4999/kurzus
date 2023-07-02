package models

type Item struct {
	ID         int64  `json:"id"`
	MenuId     int64  `json:"menu"`
	Ingerdient string `json:"ingredient"`
}
