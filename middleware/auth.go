package middleware

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github/linchao0828/messengerBot/biz/api"
	"github/linchao0828/messengerBot/biz/cache"
	"github/linchao0828/messengerBot/pkg/auth"
	"github/linchao0828/messengerBot/pkg/consts"
	"github/linchao0828/messengerBot/pkg/logger"
	"github/linchao0828/messengerBot/pkg/yerrors"
)

func Auth(anonymity bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.WithContext(c).Info("handle mw auth")
		authorization := c.GetHeader("Authorization")

		if s, ok := auth.Jwt.Verify(authorization); ok {
			if !cache.Jwt.IsBlock(c, authorization) {
				if j, err := json.Marshal(s); err == nil {
					var u auth.SessionUser
					if err := json.Unmarshal(j, &u); err == nil {
						c.Set(auth.CurrentUser, u)
						c.Set(consts.CurrentUserId, u.UserID)
						return
					}
				}
			}
		}

		session := sessions.DefaultMany(c, consts.SessionKey)
		if s := session.Get(auth.CurrentUser); s != nil {
			if u, ok := s.(auth.SessionUser); ok {
				c.Set(auth.CurrentUser, u)
				c.Set(consts.CurrentUserId, u.UserID)
				return
			}
		}
		if !anonymity {
			api.Fail(c, yerrors.StatusUnauthorized)
		}
	}
}
