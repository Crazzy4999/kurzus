package request

type ItemRequest struct {
	ID         int    `json:"id"`
	Ingerdient string `json:"ingredient"`
}
