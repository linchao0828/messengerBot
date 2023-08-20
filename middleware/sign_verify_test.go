package middleware

import (
	"fmt"
	"github/linchao0828/messengerBot/pkg/encode"
	"testing"
)

func TestSign(t *testing.T) {
	t.Run("sign", func(t *testing.T) {
		plain := "mobile=18698571125&ts=1663512944639" + "123"
		fmt.Println(encode.MD5(plain))
	})
}
