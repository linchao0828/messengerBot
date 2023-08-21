package idgen

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestNext(t *testing.T) {
	t.Run("generate id", func(t *testing.T) {
		m := make(map[int64]bool, 0)
		for i := 0; i < 1000; i++ {
			id := Next()
			println(id)
			if m[id] {
				t.FailNow()
			}
			m[id] = true
		}
	})

	t.Run("ss", func(t *testing.T) {
		println(time.Now().UnixNano() / 1e6)
	})

}

func TestUid(t *testing.T) {
	t.Run("generate uid", func(t *testing.T) {
		assert.True(t, strings.HasPrefix(strconv.Itoa(int(Uid())), "2521"))
	})
}
