package service

import (
	"github.com/gin-gonic/gin"
	"github.com/maciekmm/messenger-platform-go-sdk"
	"github/linchao0828/messengerBot/client/mp"
	"github/linchao0828/messengerBot/client/openai"
	"github/linchao0828/messengerBot/domain"
	"github/linchao0828/messengerBot/pkg/logger"
)

type messengerService struct {
}

func NewMessengerService() MessengerService {
	return &messengerService{}
}

type MessengerService interface {
	HandleMessage(ctx *gin.Context, event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) error
	HandlePostback(ctx *gin.Context, event messenger.Event, opts messenger.MessageOpts, msg messenger.Postback) error
	SendTextMessage(ctx *gin.Context, senderID, text string) error
	SendTemplateMessageWithButtons(ctx *gin.Context, senderID, text string, buttons []domain.PayloadButton) error
}

func (m messengerService) HandleMessage(ctx *gin.Context, event messenger.Event, opts messenger.MessageOpts, msg messenger.ReceivedMessage) error {
	logger.WithContext(ctx).Infof("received message: %s", msg.Text)
	messageList := openai.Cli.PackChatMessage(ctx, "I want you to act as a customer service", msg.Text, nil)
	replayMessage, _, err := openai.Cli.Chat(ctx, messageList)
	if err != nil {
		return err
	}
	return m.SendTextMessage(ctx, opts.Sender.ID, replayMessage.Content)
}

func (m messengerService) HandlePostback(ctx *gin.Context, event messenger.Event, opts messenger.MessageOpts, msg messenger.Postback) error {
	logger.WithContext(ctx).Infof("received postback, opts: %+v, msg: %+v", opts, msg)
	//根据payload来判断用户点击了哪个按钮

	//do biz

	return m.SendTextMessage(ctx, opts.Sender.ID, "Thanks for your feedback!")
}

func (m messengerService) SendTextMessage(ctx *gin.Context, senderID, text string) error {
	resp, err := mp.Cli.SendSimpleMessage(senderID, text)
	if err != nil {
		return err
	}
	logger.WithContext(ctx).Infof("sent text message to %s, text: %s, message_id: %s", senderID, text, resp.MessageID)
	return nil
}

func (m messengerService) SendTemplateMessageWithButtons(ctx *gin.Context, senderID, text string, buttons []domain.PayloadButton) error {
	resp, err := mp.Cli.SendMessage(messenger.MessageQuery{
		Recipient: messenger.Recipient{
			ID: senderID,
		},
		Message: messenger.SendMessage{
			Attachment: &messenger.Attachment{
				Type: messenger.AttachmentTypeTemplate,
				Payload: domain.Payload{
					TemplateType: domain.TemplateTypeButton,
					Text:         text,
					Buttons:      buttons,
				},
			},
		},
		MessagingType: messenger.MessagingTypeRegular,
	})
	if err != nil {
		return err
	}
	logger.WithContext(ctx).Infof("sent template message with buttons to %s, text: %s, message_id: %s", senderID, text, resp.MessageID)
	return nil
}
