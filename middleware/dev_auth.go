package middleware

import (
	"github.com/gin-gonic/gin"
	"github/linchao0828/messengerBot/biz/api"
	"github/linchao0828/messengerBot/pkg/logger"
	"github/linchao0828/messengerBot/pkg/yerrors"
)

func DevAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.WithContext(c).Info("handle mw dev_auth")
		authStr := c.GetHeader("SuperBoy")
		if authStr == "SuperMan" {
			return
		}

		logger.WithContext(c).Errorf("dev auth fail! ip:%s", c.ClientIP())
		api.Fail(c, yerrors.BadRequest)
		return
	}
}
