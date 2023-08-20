package domain

// ChatMessage open的chat消息体
type ChatMessage struct {
	Role    ChatMessageRole `json:"role"`
	Content string          `json:"content"`
}

type ChatMessageRole string

const (
	ChatMessageRole_System    ChatMessageRole = "system"
	ChatMessageRole_User      ChatMessageRole = "user"
	ChatMessageRole_Assistant ChatMessageRole = "assistant"
)
