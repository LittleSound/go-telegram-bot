package command

import (
	"fmt"
	"go-bot/context"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

// Mount 挂载 Rule 和 Call，用于导出至 controllers
type Mount struct{}

// Rule 规则
func (m Mount) Rule(ctx *context.CTX) bool {
	if !ctx.Message.IsCommand() {
		return false
	}

	if ctx.Message.Chat.Type == "private" {
		return true
	} else {
		logrus.Info(ctx.Message)

		return strings.Contains(ctx.Message.CommandWithAt(), ctx.Bot.Self.UserName)
	}
}

// Handle 处理消息
func (m Mount) Handle(ctx *context.CTX) {
	commandName := ctx.Message.Command()
	logrus.Infof("收到命令 “%s”, 参数：“%s”", commandName, ctx.Message.CommandArguments())
	for name, handler := range commands {
		if name == commandName {
			handler(ctx)
			return
		}
	}

	msg := tgbotapi.NewMessage(ctx.ChatID, "哦！我还不会这个命令……")
	msg.ReplyToMessageID = ctx.Message.MessageID
	ctx.Bot.Send(msg)
}

// Commands 命令列表
type Commands map[string]func(ctx *context.CTX)

var commands = Commands{
	"broadcast": func(ctx *context.CTX) {
		args := ctx.Message.CommandArguments()
		if ctx.Message.From.ID != 720497356 {
			msg := tgbotapi.NewMessage(ctx.ChatID, "你没有使用这个功能的权限哦")
			msg.ReplyToMessageID = ctx.Message.MessageID
			ctx.Bot.Send(msg)
			return
		}
		if args == "" {
			msg := tgbotapi.NewMessage(ctx.ChatID, "/broadcast: 广播命令\n 请在命令后跟随需要广播的文本")
			msg.ReplyToMessageID = ctx.Message.MessageID
			ctx.Bot.Send(msg)
			return
		}
	},
	"talking": func(ctx *context.CTX) {
		args := ctx.Message.CommandArguments()
		if ctx.Message.From.ID != 720497356 {
			msg := tgbotapi.NewMessage(ctx.ChatID, "你没有使用这个功能的权限哦")
			msg.ReplyToMessageID = ctx.Message.MessageID
			ctx.Bot.Send(msg)
			return
		}
		if args == "" {
			msg := tgbotapi.NewMessage(ctx.ChatID, "/talking: 说话命令\n 请在命令后跟随需要说的文本")
			msg.ReplyToMessageID = ctx.Message.MessageID
			ctx.Bot.Send(msg)
			return
		}
		msg := tgbotapi.NewMessage(-1001461459123, fmt.Sprintf("[%s %s @%s]: %s", ctx.Message.From.LastName, ctx.Message.From.FirstName, ctx.Message.From.UserName, ctx.Message.CommandArguments()))
		ctx.Bot.Send(msg)
		return
	},
}
