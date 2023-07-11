package token

import (
	"hangryAPI/internal/test/util"
	"testing"
)

func (suite *TokenServiceTestSuite) TestValidateAccessToken() {
	tokenString, _ := suite.tokenService.GenerateAccessToken(userID)
	invalidTokenString := tokenString + "asd"

	testCases := []util.TestCaseValidate{
		{
			Name:         "Valid token",
			TokenString:  tokenString,
			WantError:    false,
			WantErrorMsg: "",
			WantID:       userID,
		},
		{
			Name:         "Invalid token",
			TokenString:  invalidTokenString,
			WantError:    true,
			WantErrorMsg: "signature is invalid",
			WantID:       0,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.Name, func(t *testing.T) {
			gotClaims, err := suite.tokenService.ValidateAccessToken(testCase.TokenString)

			util.AssertTokenValidationResult(t, testCase, gotClaims, err)
		})
	}
}

func (suite *TokenServiceTestSuite) TestValidaterefreshToken() {
	tokenString, _ := suite.tokenService.GenerateRefreshToken(userID)
	invalidTokenString := tokenString + "asd"

	testCases := []util.TestCaseValidate{
		{
			Name:         "Valid token",
			TokenString:  tokenString,
			WantError:    false,
			WantErrorMsg: "",
			WantID:       userID,
		},
		{
			Name:         "Invalid token",
			TokenString:  invalidTokenString,
			WantError:    true,
			WantErrorMsg: "signature is invalid",
			WantID:       0,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.Name, func(t *testing.T) {
			gotClaims, err := suite.tokenService.ValidateRefreshToken(testCase.TokenString)

			util.AssertTokenValidationResult(t, testCase, gotClaims, err)
		})
	}
}

func (suite *TokenServiceTestSuite) TestValidateResetToken() {
	tokenString, _ := suite.tokenService.GenerateResetToken(userID)
	invalidTokenString := tokenString + "asd"

	testCases := []util.TestCaseValidate{
		{
			Name:         "Valid token",
			TokenString:  tokenString,
			WantError:    false,
			WantErrorMsg: "",
			WantID:       userID,
		},
		{
			Name:         "Invalid token",
			TokenString:  invalidTokenString,
			WantError:    true,
			WantErrorMsg: "signature is invalid",
			WantID:       0,
		},
	}

	for _, testCase := range testCases {
		suite.T().Run(testCase.Name, func(t *testing.T) {
			gotClaims, err := suite.tokenService.ValidateResetToken(testCase.TokenString)

			util.AssertTokenValidationResult(t, testCase, gotClaims, err)
		})
	}
}
