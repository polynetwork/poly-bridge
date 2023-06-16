package common

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"poly-bridge/conf"
)

var bot *tgbotapi.BotAPI

func TgBotInit() {
	var err error
	if len(conf.GlobalConfig.BotConfig.TgBotApiToken) == 0 {
		panic("TgBotApiToken is empty")
	}

	bot, err = tgbotapi.NewBotAPI(conf.GlobalConfig.BotConfig.TgBotApiToken)
	if err != nil {
		panic(fmt.Sprintf("tg bot init failed:%s", err))
	}
}

func SendTgBotMessage(msg tgbotapi.MessageConfig) (tgbotapi.Message, error) {
	return bot.Send(msg)
}
