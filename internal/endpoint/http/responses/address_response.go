package responses

type AddressResponse struct {
	CityName    string `json:"cityName"`
	StreetName  string `json:"streetName"`
	HouseNumber int    `json:"houseNumber"`
	FloorNumber int    `json:"floorNumber"`
	Apartment   string `json:"apartment"`
}
