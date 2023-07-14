package responses

type OrderMenusCollectionResponse struct {
	OrderMenus []*OrderMenuResponse `json:"orderMenus"`
}
