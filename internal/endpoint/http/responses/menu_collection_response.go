package responses

type MenuCollectionResponse struct {
	Menus []*MenuResponse `json:"menus"`
}
