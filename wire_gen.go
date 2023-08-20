// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github/linchao0828/messengerBot/biz/api"
	"github/linchao0828/messengerBot/biz/router"
	"github/linchao0828/messengerBot/biz/service"
)

import (
	_ "github/linchao0828/messengerBot/docs"
)

// Injectors from wire.go:

func BuildController() router.Controller {
	messengerService := service.NewMessengerService()
	orderService := service.NewOrderService(messengerService)
	messenger := api.Messenger{
		MessengerService: messengerService,
		OrderService:     orderService,
	}
	controller := router.Controller{
		MessengerAPI: messenger,
	}
	return controller
}

// wire.go:

var ServiceSet = wire.NewSet(service.NewMessengerService, service.NewOrderService)

var ApiSet = wire.NewSet(wire.Struct(new(api.Messenger), "*"))