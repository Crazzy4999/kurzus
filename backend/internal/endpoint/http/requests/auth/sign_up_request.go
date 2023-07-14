package request

type SignUpRequest struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	PasswordAgain string `json:"passwordAgain"`
}
