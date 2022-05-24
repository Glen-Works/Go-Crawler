package service

import (
	"crawler/project/internal/utils"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramRobot struct {
	groupId string
	token   string
}

func NewTelegramRobotService() *TelegramRobot {
	return &TelegramRobot{
		groupId: "TELEGRAM_GROUP_ID",
		token:   "TELEGRAM_TOKEN",
	}
}

func (tr *TelegramRobot) sendMsg(msg string) {
	var err error
	bot, err := tgbotapi.NewBotAPI(utils.GetEnvData(tr.token))
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = false

	groupId, err := strconv.ParseInt(utils.GetEnvData(tr.groupId), 10, 64)
	if err != nil {
		log.Fatal("Telegram group id err")
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
