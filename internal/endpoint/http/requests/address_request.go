package request

type AddressRequest struct {
	City        string `json:"city"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	ZipCode     string `json:"zipCode"`
	FloorNumber string `json:"floorNumber"`
	Apartment   string `json:"apartment"`
}
