package models

const (
	CREATED = iota + 1
	DELIVERING
	DONE
)

type OrderStatus struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}
