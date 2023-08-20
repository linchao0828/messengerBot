package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/linchao0828/messengerBot/biz/api"
	"github/linchao0828/messengerBot/pkg/encode"
	"github/linchao0828/messengerBot/pkg/logger"
	"github/linchao0828/messengerBot/pkg/yerrors"
	"sort"
	"strings"
)

const (
	signQueryKey = "sign"
)

func SignVerify(slat string) gin.HandlerFunc {
	return func(c *gin.Context) {

		keys := make([]string, 0)
		for k := range c.Request.URL.Query() {
			keys = append(keys, k)
		}

		sort.Strings(keys)
		queries := make([]string, 0)
		for _, k := range keys {
			if k != signQueryKey {
				queries = append(queries, fmt.Sprintf("%s=%s", k, c.Query(k)))
			}
		}

		plain := strings.Join(queries, "&") + slat
		expected := encode.MD5(plain)
		actual := c.Query(signQueryKey)
		if actual != expected {
			logger.WithContext(c).Errorf("verify sign fail! expected:%s, actual:%s", expected, actual)
			api.Fail(c, yerrors.BadRequest)
			return
		}
	}
}
