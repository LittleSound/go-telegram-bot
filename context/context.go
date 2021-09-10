package context

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// CTX 消息上下文
type CTX struct {
	Bot     *tgbotapi.BotAPI
	Message *tgbotapi.Message
	ChatID  int64
}

// NewCTX 新建 上下文
func NewCTX(Bot *tgbotapi.BotAPI, message *tgbotapi.Message) *CTX {
	return &CTX{
		Bot:     Bot,
		Message: message,
		ChatID:  message.Chat.ID,
	}
}
