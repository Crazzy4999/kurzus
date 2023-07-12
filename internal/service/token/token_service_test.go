package token

import (
	"fmt"
	config "hangryAPI/configs"
	"hangryAPI/internal/test/util"
	testcases "hangryAPI/internal/test/util"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const userID = 1

type TokenServiceTestSuite struct {
	suite.Suite
	cfg          *config.Config
	tokenService *TokenService
	testSrv      *httptest.Server
}

func (suite *TokenServiceTestSuite) SetupSuite() {
	suite.cfg = config.NewConfig("../../../configs/.env.testing")
	suite.tokenService = NewTokenService(suite.cfg)
}

func (suite *TokenServiceTestSuite) TearDownSuite() {}

func (suite *TokenServiceTestSuite) SetupTest() {}

func (suite *TokenServiceTestSuite) TearDownTest() {}

func TestTokenServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TokenServiceTestSuite))
}

func (suite *TokenServiceTestSuite) TestGetTokenFromBearerString() {
	testCases := []testcases.TestCaseGetBearerToken{
		{
			Name:         "Get token successfully",
			BearerString: "Bearer test_token",
			Want:         "test_token",
		},
		{
			Name:         "Return empty string if Bearer string is empty",
			BearerString: "",
			Want:         "",
		},
		{
			Name:         "Return empty string if Bearer prefix is invalid",
			BearerString: "Bear test_token",
			Want:         "",
		},
		{
			Name:         "Return empty string if there is no token string",
			BearerString: "Bearer ",
			Want:         "",
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.Name, func(t *testing.T) {
			got := GetTokenFromBearerString(testCase.BearerString)
			assert.Equal(suite.T(), testCase.Want, got)
		})
	}
}

func (suite *TokenServiceTestSuite) TestGetClaims() {
	testCases := []testcases.TestCaseHandler{
		{
			Name: "Get claims successfully",
			Request: util.Request{
				Method:    http.MethodGet,
				URL:       "/profile",
				AuthToken: suite.cfg.AccessSecret,
			},
			Want: util.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "1",
			},
		},
		{
			Name: "Get claims failed",
			Request: util.Request{
				Method:    http.MethodGet,
				URL:       "/profile",
				AuthToken: "",
			},
			Want: util.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   "",
			},
		},
	}

	for _, testCase := range testCases {
		request := testcases.TestRequestFactory(testCase, testCase.Request.URL)

		accessToken, err := suite.tokenService.GenerateAccessToken(userID)
		assert.NoError(suite.T(), err)

		fmt.Println(accessToken)

		request.Header.Set("Authorization", "Bearer "+accessToken)

		claims, err := suite.tokenService.GetClaims(request, testCase.Request.AuthToken)

		if testCase.Request.AuthToken == "" {
			assert.Error(suite.T(), err)
		} else {
			id, err := strconv.Atoi(testCase.Want.BodyPart)
			assert.NoError(suite.T(), err)
			assert.Equal(suite.T(), id, claims.ID)
		}
	}

}

func (suite *TokenServiceTestSuite) TestGetClaimsFromContext() {

}
