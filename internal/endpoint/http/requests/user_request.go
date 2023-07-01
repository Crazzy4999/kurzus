package request

type UserRequest struct {
	Address   int64  `json:"address"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
