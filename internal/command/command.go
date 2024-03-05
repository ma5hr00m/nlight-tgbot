package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"nlight-tgbot/internal/config"
	"nlight-tgbot/internal/controller/juejin"
)

// 定义指令和描述的映射
var commands = map[string]string{
	"help":    "Show available commands",
	"ping":    "Ping command",
	"checkin": "Juejin Check-in command",
}

func HandleCommands(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Command() {
	case "help":
		// 构建包含所有指令和描述的消息
		var helpText string
		helpText += "你好，我是 NLight Bot！目前支持以下命令：\n"
		for cmd, desc := range commands {
			helpText += "/" + cmd + ": " + desc + "\n"
		}
		msg.Text = helpText
	case "ping":
		msg.Text = "pong!"
	case "checkin":
		msg.Text = juejin.DailyCheckIn(config.Juejin.Uuid, config.Juejin.SessionId)
	default:
		msg.Text = "I don't know that command."
	}

	return msg
}
