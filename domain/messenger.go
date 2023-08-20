package domain

type TemplateType string

const (
	TemplateTypeButton TemplateType = "button"
)

type PayloadButtonType string

const (
	ButtonTypePostback PayloadButtonType = "postback"
)

type Payload struct {
	TemplateType TemplateType    `json:"template_type"`
	Text         string          `json:"text"`
	Buttons      []PayloadButton `json:"buttons"`
}

type PayloadButton struct {
	Type    PayloadButtonType `json:"type"`
	URL     string            `json:"url"`
	Payload string            `json:"payload"`
	Title   string            `json:"title"`
}

type MockOrderSuccessReq struct {
	OrderID           string `json:"order_id" binding:"required"`
	MessengerSenderID string `json:"messenger_sender_id" binding:"required"`
}
