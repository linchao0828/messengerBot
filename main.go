package main

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github/linchao0828/messengerBot/client/mp"
	"github/linchao0828/messengerBot/client/openai"
	"github/linchao0828/messengerBot/conf"
	_ "github/linchao0828/messengerBot/docs"
	"github/linchao0828/messengerBot/middleware"
	"github/linchao0828/messengerBot/pkg/auth"
	"github/linchao0828/messengerBot/pkg/logger"
	"log"
)

// @title       messengerBot
// @version     1.0
// @description This is messengerBot server.
func main() {
	var err error
	conf.LoadConf()
	logger.Init(conf.Config.LogLevel)

	r := gin.New()
	// mw
	//path := "./logs/runtime.log"
	//f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	//r.Use(gin.LoggerWithWriter(f, "/"))
	//r.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/"}}))
	r.Use(gin.Recovery())
	r.Use(requestid.New())
	pprof.Register(r)

	// session auth
	gob.Register(auth.SessionUser{})
	//gob.Register(auth.SessionManager{})
	//store, err := redis.NewStore(10, "tcp", conf.Config.Redis.Addr, conf.Config.Redis.Pwd, []byte("secret"))
	//if err != nil {
	//	panic(err)
	//}
	//r.Use(sessions.SessionsMany([]string{consts.SessionKey, consts.ManagerSessionKey}, store))

	// log id
	r.Use(middleware.Access())
	//dal.Init(conf.Config.DSN)
	//infra.InitRedis(conf.Config.Redis.Addr, conf.Config.Redis.Pwd, conf.Config.Redis.DB)

	// clients
	openai.Init(conf.Config.OpenAI.AuthKey, conf.Config.OpenAI.ChatCompletionsUrl, conf.Config.OpenAI.Model, conf.Config.OpenAI.Temperature, conf.Config.OpenAI.MaxTokens, conf.Config.HttpProxy.Url)
	mp.Init(conf.Config.Messenger.AccessToken, conf.Config.Messenger.VerifyToken)

	// router
	BuildController().Register(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// cron
	//if conf.Config.Env != "dev" {
	//	BuildCron().Start()
	//}

	if err = r.Run(fmt.Sprintf(":%d", conf.Config.Port)); err != nil {
		log.Fatal(err)
	}
}
