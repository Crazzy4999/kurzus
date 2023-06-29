package request

type ResetPasswordRequest struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	PasswordAgain string `json:"passwordAgain"`
}
