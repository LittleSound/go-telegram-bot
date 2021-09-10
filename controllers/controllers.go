package controllers

import (
	"go-bot/context"
	"go-bot/controllers/command"
	"go-bot/controllers/repeater"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

// Controller 调度类型
type Controller interface {
	// 判断当前消息是否被该程序处理
	Rule(ctx *context.CTX) bool
	// 消息处理程序
	Handle(ctx *context.CTX)
}

// Controllers 调度列表
type Controllers map[string]Controller

// 挂载点
var controllers = Controllers{
	"repeater": repeater.Mount{},
	"command":  command.Mount{},
}

// Sorters 将消息分配到对应的处理程序
func Sorters(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	ctx := context.NewCTX(bot, message)
	for name, controller := range controllers {
		if controller.Rule(ctx) {
			logrus.Infof("被：‘%s’ 调度器接收", name)
			go controller.Handle(ctx)
			return
		}
	}
	logrus.Info("未处理此消息")
}

// WaitReply 等待回复
func WaitReply(ctx context.CTX) {

}
