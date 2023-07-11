package util

import (
	"net/http"
	"strings"
)

type TestCaseGetBearerToken struct {
	Name         string
	BearerString string
	Want         string
}

type TestCaseValidate struct {
	Name         string
	TokenString  string
	WantError    bool
	WantErrorMsg string
	WantID       int
}

type Request struct {
	Method    string
	URL       string
	AuthToken string
}

type ExpectedResponse struct {
	StatusCode int
	BodyPart   string
}

type TestCaseHandler struct {
	Name        string
	Request     Request
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
	Want        ExpectedResponse
}

func TestRequestFactory(test TestCaseHandler, testURL string) *http.Request {
	request, _ := http.NewRequest(test.Request.Method, testURL+test.Request.URL, strings.NewReader(""))

	if test.Request.AuthToken == "" {
		request.Header.Set("Authorization", "Bearer "+test.Request.AuthToken)
	}

	return request
}
