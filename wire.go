//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github/linchao0828/messengerBot/biz/api"
	"github/linchao0828/messengerBot/biz/router"
	"github/linchao0828/messengerBot/biz/service"
)

var ServiceSet = wire.NewSet(
	service.NewMessengerService,
	service.NewOrderService,
)

var ApiSet = wire.NewSet(
	wire.Struct(new(api.Messenger), "*"),
)

func BuildController() router.Controller {
	panic(wire.Build(
		wire.Struct(new(router.Controller), "*"),
		ApiSet,
		ServiceSet,
	))
}

//func BuildCron() cron.Cron {
//	panic(wire.Build(
//		wire.Struct(new(cron.Cron), "*"),
//		ServiceSet,
//	))
//}
