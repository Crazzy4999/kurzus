package responses

type DriverCollectionResponse struct {
	Drivers []*DriverResponse `json:"drivers"`
}
