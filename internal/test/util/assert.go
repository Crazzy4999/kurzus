package util

import (
	"hangryAPI/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertTokenValidationResult(t *testing.T, testCase TestCaseValidate, gotClaims *util.JwtCustomClaims, err error) {
	t.Helper()

	if testCase.WantError {
		assert.Error(t, err)
		assert.Contains(t, err.Error(), testCase.WantErrorMsg)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, testCase.WantID, gotClaims.ID)
	}
}
