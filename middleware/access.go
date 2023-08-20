package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github/linchao0828/messengerBot/pkg/consts"
)

func Access() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(consts.LogId, uuid.New().String())
	}
}
