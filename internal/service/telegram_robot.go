package service

import (
	"crawler/project/internal/utils"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramRobot struct {
}

func NewTelegramRobotService() *TelegramRobot {
	return &TelegramRobot{}
}

func (tr *TelegramRobot) SendMsg(msg string, cc *utils.CrawlerConfig) {

	bot, err := tgbotapi.NewBotAPI(cc.TelegramToken)
	if err != nil {
		log.Panicf("Telegram token err:%s", err)
	}
	bot.Debug = false

	groupId, err := strconv.ParseInt(cc.TelegramGroupId, 10, 64)
	if err != nil {
		log.Panicf("Telegram group id err:%s", err)
	}

	NewMsg := tgbotapi.NewMessage(groupId, msg)
	// NewMsg.ParseMode = tgbotapi.ModeHTML //傳送html格式的訊息

	_, err = bot.Send(NewMsg)
	if err == nil {
		log.Printf("Send telegram message success")
	} else {
		log.Printf("Send telegram message error")
	}
}
