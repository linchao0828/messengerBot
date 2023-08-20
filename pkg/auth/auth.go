package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github/linchao0828/messengerBot/conf"
	"github/linchao0828/messengerBot/pkg/consts"
	"net/http"
	"strings"
)

const (
	CurrentUser        = "CURRENT_USER"
	CurrentManagerUser = "CURRENT_MANAGER_USER"
)

type SessionUser struct {
	UserID int64
	Email  string
}

type SessionManager struct {
	Id       int64
	Username string
}

func Save(c *gin.Context, s SessionUser) {
	session := sessions.DefaultMany(c, consts.SessionKey)
	session.Set(CurrentUser, s)
	_ = session.Save()
}

func Clear(c *gin.Context) {
	session := sessions.DefaultMany(c, consts.SessionKey)
	session.Delete(CurrentUser)
	session.Options(sessions.Options{Path: "/", MaxAge: -1})

	_ = session.Save()
}

func SaveManager(c *gin.Context, s SessionManager) {
	session := sessions.DefaultMany(c, consts.ManagerSessionKey)
	session.Set(CurrentManagerUser, s)
	domain := c.GetHeader("referer")
	if domain == "" {
		domain = conf.Config.Cookie.Domain
	}
	domain = strings.TrimSuffix(domain, "/")
	domain = strings.TrimPrefix(domain, "http://")
	domain = strings.TrimPrefix(domain, "https://")

	session.Options(sessions.Options{
		Path:     "/",
		Domain:   domain,
		MaxAge:   2592000,
		Secure:   false,
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
	})
	_ = session.Save()
}

func ClearManager(c *gin.Context) {
	session := sessions.DefaultMany(c, consts.ManagerSessionKey)
	session.Delete(CurrentManagerUser)
	domain := c.GetHeader("referer")
	if domain == "" {
		domain = conf.Config.Cookie.Domain
	}
	domain = strings.TrimSuffix(domain, "/")
	domain = strings.TrimPrefix(domain, "http://")
	domain = strings.TrimPrefix(domain, "https://")

	session.Options(sessions.Options{
		Path:     "/",
		Domain:   domain,
		MaxAge:   -1,
		Secure:   false,
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
	})

	_ = session.Save()
}
