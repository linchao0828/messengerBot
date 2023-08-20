package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaskAddress(t *testing.T) {
	t.Run("mask address", func(t *testing.T) {
		res := MaskAddress("cfxtest:adw0aj3c3j5p6fzz096gwgby6wnph6g82pf166jf8g")
		assert.Equal(t, "cfxtest:adw0...jf8g", res)
	})
}
