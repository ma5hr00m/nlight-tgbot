package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"nlight-tgbot/internal/config"
	"nlight-tgbot/internal/controller/blog"
	"nlight-tgbot/internal/controller/juejin"
)

// 定义指令和描述的映射
var commands = map[string]string{
	"help":    "Show available commands",
	"ping":    "Ping command",
	"checkin": "Juejin Check-in command",
	"blog":    "Detect blog status",
}

func HandleCommands(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Command() {
	case "help":
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
	case "blog":
		msg.Text = blog.BlogSurvivalDetect(config.Blog.Title, config.Blog.URL)
	default:
		msg.Text = "I don't know that command."
	}

	return msg
}
