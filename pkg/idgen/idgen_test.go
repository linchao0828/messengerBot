package idgen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestNexta(t *testing.T) {
	name := "王禹衡"

	//s := []rune(name)[0]
	println(string(name[2:]))
}
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

func TestCid(t *testing.T) {
	t.Run("generate cid", func(t *testing.T) {
		for i := 0; i < 500; i++ {
			fmt.Println(CharacterID())
			time.Sleep(10 * time.Millisecond)
		}
	})
}
