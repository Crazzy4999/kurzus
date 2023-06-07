package model

type StatusType string

const (
	CREATED     StatusType = "created"
	IN_PROGRESS StatusType = "in_progress"
	DONE        StatusType = "done"
)

type Order struct {
	OrderId    int `json:"order_id"`
	SupplierId int `json:"supplier_id"`
	UserId     int `json:"user_id"`
	Address    struct {
		Street   string `json:"street"`
		Building string `json:"building"`
		Apt      int    `json:"apt"`
	} `json:"address"`
	Status StatusType `json:"status"`
}
