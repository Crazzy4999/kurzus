package request

type AddressRequest struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userID"`
	IsActive    bool   `json:"isActive"`
	City        string `json:"city"`
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	ZipCode     string `json:"zipCode"`
	FloorNumber string `json:"floorNumber"`
	Apartment   string `json:"apartment"`
}
