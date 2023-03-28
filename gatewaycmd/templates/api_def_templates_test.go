package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFullProtectedApiDefinition(t *testing.T) {
	got, err := createFullProtectedAPIDefinition()

	assert.Nil(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, "4a77d8bfe76f41ad7ae5875b2259df3f", *got.ApiId)
	assert.Equal(t, "Authorization", *got.Auth.AuthHeaderName)
	assert.Equal(t, float64(-1), *got.HmacAllowedClockSkew)
}
