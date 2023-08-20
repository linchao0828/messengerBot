package router

import (
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"github/linchao0828/messengerBot/biz/api"
	"github/linchao0828/messengerBot/middleware"
	"net/http"
)

type Controller struct {
	MessengerAPI api.Messenger
}

func (c Controller) Register(r *gin.Engine) {
	// health check
	r.HEAD("", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
		return
	})
	r.Use(middleware.Cors())

	v := r.Group("/api/v1")
	v.Use(gindump.DumpWithOptions(true, true, true, false, true, nil))
	{
		mp := v.Group("/mp")
		{
			mp.Any("/webhook", c.MessengerAPI.Webhook)
			mp.Any("/mock/order_success", c.MessengerAPI.MockOrderSuccess)
		}
	}
	c.registerAdmin(r)
}

func (c Controller) registerAdmin(r *gin.Engine) {
	v := r.Group("/api/v1/admin")
	v.Use(gindump.DumpWithOptions(true, true, true, false, true, nil))
	{
	}
}
