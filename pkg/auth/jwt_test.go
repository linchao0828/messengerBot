package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	t.Run("should codec", func(t *testing.T) {
		token, err := Jwt.Create("fake-id")
		assert.Nil(t, err)
		res, ok := Jwt.Verify(token)
		assert.True(t, ok)
		assert.Equal(t, "fake-id", res)
	})
	t.Run("should fail when token invalid", func(t *testing.T) {
		res, ok := Jwt.Verify("abc")
		assert.False(t, ok)
		assert.Empty(t, res)
	})
}
