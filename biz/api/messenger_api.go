package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maciekmm/messenger-platform-go-sdk"
	"github/linchao0828/messengerBot/biz/service"
	"github/linchao0828/messengerBot/client/mp"
	"github/linchao0828/messengerBot/domain"
	"github/linchao0828/messengerBot/pkg/logger"
	"github/linchao0828/messengerBot/pkg/yerrors"
)

type Messenger struct {
	MessengerService service.MessengerService
	OrderService     service.OrderService
}

// Webhook godoc
// @Summary webhook
// @Tags    messenger
// @Accept  json
// @Produce json
// @Success 200     {object} BaseResp
// @Router  /api/v1/mp/webhook [post]
// @Router  /api/v1/mp/webhook [get]
func (m *Messenger) Webhook(ctx *gin.Context) {
	cli := mp.Cli
	cli.MessageReceived = func(event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) {
		err := m.MessengerService.HandleMessage(ctx, event, opts, msg)
		if err != nil {
			logger.WithContext(ctx).WithError(err).Errorf("MessageReceived error")
		}
		return
	}
	cli.Postback = func(event messenger.Event, opts messenger.MessageOpts, msg messenger.Postback) {
		err := m.MessengerService.HandlePostback(ctx, event, opts, msg)
		if err != nil {
			logger.WithContext(ctx).WithError(err).Errorf("Postback error")
		}
		return
	}
	cli.Handler(ctx.Writer, ctx.Request)
}

// MockOrderSuccess godoc
// @Summary mock order success
// @Tags    messenger
// @Accept  json
// @Produce json
// @Param   request body domain.MockOrderSuccessReq true "mock order success request"
// @Success 200     {object} BaseResp
// @Router  /api/v1/mp/mock/order_success [post]
func (m *Messenger) MockOrderSuccess(ctx *gin.Context) {
	var req domain.MockOrderSuccessReq
	if err := ctx.BindJSON(&req); err != nil {
		logger.WithContext(ctx).WithError(err).Error("bind req json err")
		Fail(ctx, yerrors.BadRequest)
		return
	}

	err := m.OrderService.MockOrderSuccess(ctx, req)
	if err != nil {
		Fail(ctx, err)
		return
	}

	OK(ctx, nil)
}
