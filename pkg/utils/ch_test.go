package utils

import (
	"fmt"
	"testing"
)

func TestCh(t *testing.T) {
	t.Run("mark ch name", func(t *testing.T) {
		s := "我是张三"
		fmt.Println(Mark.Mark(s, 0, CH.Len(s)-1))
		fmt.Println(Mark.Mark(s, 1, CH.Len(s)-1))
	})
}
