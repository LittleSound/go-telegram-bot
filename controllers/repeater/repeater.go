package repeater

import (
	"go-bot/context"
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

// Mount 挂载 Rule 和 Call，用于导出至 controllers
type Mount struct{}

// Rule 规则
func (m Mount) Rule(ctx *context.CTX) bool {
	probability := rand.Intn(100)
	logrus.Infof("是否复读：%v < 2", probability)
	return probability < 2
}

// Handle 执行复读
func (m Mount) Handle(ctx *context.CTX) {
	for _, handleFunc := range replyHandler {
		if handleFunc(ctx) {
			return
		}
	}
}

// 回复类型消息处理，如果不能处理返回 false，如果可以处理返回 true
var replyHandler = map[string]func(ctx *context.CTX) bool{
	"text": func(ctx *context.CTX) bool {
		if ctx.Message.Text == "" {
			return false
		}

		msg := tgbotapi.NewMessage(ctx.ChatID, ctx.Message.Text)
		ctx.Bot.Send(msg)
		return true
	},
	"sticker": func(ctx *context.CTX) bool {
		if ctx.Message.Sticker == nil {
			return false
		}

		msg := tgbotapi.NewStickerShare(ctx.ChatID, ctx.Message.Sticker.FileID)
		ctx.Bot.Send(msg)
		return true
	},
}
