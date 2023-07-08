package util_test

type TestCaseGetBearerToken struct {
	Name         string
	BearerString string
	Want         string
}

type TestCaseValidate struct {
	Name         string
	AccessToken  string
	WantError    bool
	WantErrorMsg string
	WantID       int
}
