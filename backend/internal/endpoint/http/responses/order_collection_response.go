package responses

type OrdersCollectionResponse struct {
	Orders []*OrderResponse `json:"orders"`
}
