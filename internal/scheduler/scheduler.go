package scheduler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron"
	"nlight-tgbot/internal/config"
	"nlight-tgbot/internal/controller/juejin"
)

func StartScheduler(bot *tgbotapi.BotAPI) {
	c := cron.New()

	// 在每天上午10:30执行 DailyCheckIn 函数
	c.AddFunc("0 30 10 * * *", func() {
		message := juejin.DailyCheckIn(config.Juejin.Uuid, config.Juejin.SessionId)
		SendMessageToGroup(bot, message)
	})

	c.Start()
}

func SendMessageToGroup(bot *tgbotapi.BotAPI, message string) {
	msg := tgbotapi.NewMessage(config.Bot.GroupId, message)
	_, err := bot.Send(msg)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}
