package token

import (
	"hangryAPI/internal/util"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func (suite *TokenServiceTestSuite) TestGenerateAccessToken() {
	accessString, err := suite.tokenService.GenerateAccessToken(userID)
	assert.NoError(suite.T(), err)

	token, err := jwt.ParseWithClaims(accessString, &util.JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(suite.cfg.AccessSecret), nil
	})
	assert.NoError(suite.T(), err)

	claims, ok := token.Claims.(*util.JwtCustomClaims)
	assert.True(suite.T(), ok)
	assert.True(suite.T(), token.Valid)

	got := claims.ID
	assert.Equal(suite.T(), userID, got)

	expireTime := time.Unix(claims.ExpiresAt.Unix(), 0)
	assert.WithinDuration(suite.T(), time.Now().Add(time.Minute*time.Duration(suite.cfg.AccessLifetime)), expireTime, time.Second)
}

func (suite *TokenServiceTestSuite) TestGenerateRefreshToken() {
	refreshString, err := suite.tokenService.GenerateRefreshToken(userID)
	assert.NoError(suite.T(), err)

	token, err := jwt.ParseWithClaims(refreshString, &util.JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(suite.cfg.RefreshSecret), nil
	})
	assert.NoError(suite.T(), err)

	claims, ok := token.Claims.(*util.JwtCustomClaims)
	assert.True(suite.T(), ok)
	assert.True(suite.T(), token.Valid)

	got := claims.ID
	assert.Equal(suite.T(), userID, got)

	expireTime := time.Unix(claims.ExpiresAt.Unix(), 0)
	assert.WithinDuration(suite.T(), time.Now().Add(time.Minute*time.Duration(suite.cfg.RefreshLifetime)), expireTime, time.Second)
}

func (suite *TokenServiceTestSuite) TestGenerateResetToken() {
	resetString, err := suite.tokenService.GenerateResetToken(userID)
	assert.NoError(suite.T(), err)

	token, err := jwt.ParseWithClaims(resetString, &util.JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(suite.cfg.ResetSecret), nil
	})
	assert.NoError(suite.T(), err)

	claims, ok := token.Claims.(*util.JwtCustomClaims)
	assert.True(suite.T(), ok)
	assert.True(suite.T(), token.Valid)

	got := claims.ID
	assert.Equal(suite.T(), userID, got)

	expireTime := time.Unix(claims.ExpiresAt.Unix(), 0)
	assert.WithinDuration(suite.T(), time.Now().Add(time.Minute*time.Duration(suite.cfg.ResetLifetime)), expireTime, time.Second)
}
