package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/linchao0828/messengerBot/client/mp"
	"github/linchao0828/messengerBot/domain"
)

type orderService struct {
	messengerService MessengerService
}

func NewOrderService(messengerService MessengerService) OrderService {
	return &orderService{
		messengerService: messengerService,
	}
}

type OrderService interface {
	MockOrderSuccess(ctx *gin.Context, req domain.MockOrderSuccessReq) error
}

func (o orderService) MockOrderSuccess(ctx *gin.Context, req domain.MockOrderSuccessReq) error {
	profile, err := mp.Cli.GetProfile(req.MessengerSenderID)
	if err != nil {
		return err
	}
	//send template message to user
	text := fmt.Sprintf("Dear %s %s, your order: %s is completed. Please use the button below to give us feedback. Have a nice day!",
		profile.FirstName, profile.LastName, req.OrderID)
	buttons := []domain.PayloadButton{
		{
			Type:    domain.ButtonTypePostback,
			Title:   "Good",
			Payload: fmt.Sprintf("order_feedback_g_%s", req.OrderID),
		},
		{
			Type:    domain.ButtonTypePostback,
			Title:   "Normal",
			Payload: fmt.Sprintf("order_feedback_n_%s", req.OrderID),
		},
		{
			Type:    domain.ButtonTypePostback,
			Title:   "Bad",
			Payload: fmt.Sprintf("order_feedback_b_%s", req.OrderID),
		},
	}
	return o.messengerService.SendTemplateMessageWithButtons(ctx, req.MessengerSenderID, text, buttons)
}
