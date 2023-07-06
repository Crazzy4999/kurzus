package request

type UserRequest struct {
	Address   int    `json:"address"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
