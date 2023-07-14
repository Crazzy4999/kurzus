package responses

type CategoryCollectionResponse struct {
	Categories []*CategoryResponse `json:"categories"`
}
