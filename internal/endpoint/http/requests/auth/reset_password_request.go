package request

type ResetPasswordRequest struct {
	Password      string `json:"password"`
	PasswordAgain string `json:"passwordAgain"`
}
