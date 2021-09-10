package main

import (
	"fmt"
	"go-bot/assertion"
	"go-bot/config"
	"go-bot/controllers"
	"log"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

// AppConfig 程序配置
var AppConfig config.Config

func main() {
	fmt.Printf("%+v\n", config.Config{})
	AppConfig = config.InitConfig("./config/config.yaml")

	bot, err := tgbotapi.NewBotAPI(AppConfig.BotToken)
	assertion.Panic(err)
	goLiveTime := time.Now().Unix()
	bot.Debug = AppConfig.Debug
	log.Printf("授权账户 %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	assertion.Error(err)

	rand.Seed(time.Now().UnixNano())
	for update := range updates {
		if update.Message == nil {
			continue
		}
		logrus.Infof("消息时间：%v <<< %v", update.Message.Date, int(goLiveTime))
		if AppConfig.IgnoreThePrelaunchMessages && update.Message.Date < int(goLiveTime) {
			continue
		}
		logrus.Infof("收到消息对象：%+v\n", update.Message)
		logStr := fmt.Sprintf("Chat: %v [%s (ID: %v)]", update.Message.Chat.ID, update.Message.From.UserName, update.Message.From.ID)

		message := update.Message
		if message.Text != "" {
			logStr += fmt.Sprintf("说了：%s\n", message.Text)
		} else if message.Sticker != nil {
			logStr += fmt.Sprintf("贴纸：%s: %s %s", message.Sticker.SetName, message.Sticker.Emoji, message.Sticker.FileID)
		} else {
			logStr += "未知的消息类型"
		}

		logrus.Info(logStr)
		controllers.Sorters(bot, message)
	}
}
