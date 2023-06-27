package responses

type AddressResponse struct {
	Country     string `json:"country"`
	Region      string `json:"region"`
	CityName    string `json:"cityName"`
	StreetName  string `json:"streetName"`
	HouseNumber int    `json:"houseNumber"`
	FloorNumber int    `json:"floorNumber"`
}
