package responses

type AddressCollectionResponse struct {
	Addresses []*AddressResponse `json:"addresses"`
}
