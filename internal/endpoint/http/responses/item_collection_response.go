package responses

type ItemCollectionResponse struct {
	Items []*ItemResponse `json:"items"`
}
