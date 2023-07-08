package token

import (
	config "hangryAPI/configs"

	"github.com/stretchr/testify/suite"
)

type TokenServiceTestSuite struct {
	suite.Suite
	cfg          *config.Config
	tokenService *TokenService
}

func (suite *TokenServiceTestSuite) SetupSuite() {
	suite.cfg = &config.Config{
		AccessSecret:    "access_secret",
		AccessLifetime:  1,
		RefreshSecret:   "refresh_secret",
		RefreshLifetime: 1,
	}
	suite.tokenService = NewTokenService(suite.cfg)
}

func (suite *TokenServiceTestSuite) TearDownSuite() {}

func (suite *TokenServiceTestSuite) SetupTest() {}

func (suite *TokenServiceTestSuite) TearDownTest() {}

func (suite *TokenServiceTestSuite) TestGetTokenFromBearerString() {

}

func (suite *TokenServiceTestSuite) TestGetClaims() {

}

func (suite *TokenServiceTestSuite) TestGetClaimsFromContext() {

}
