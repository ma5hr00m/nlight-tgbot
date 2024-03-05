package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"nlight-tgbot/internal/command"
	"nlight-tgbot/internal/config"
	"nlight-tgbot/internal/scheduler"
)

func Run() {
	bot, err := tgbotapi.NewBotAPI(config.Bot.BotToken)
	if err != nil {
		panic(err)
	}
	bot.Debug = config.Bot.Debug
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = config.Bot.Timeout
	updates := bot.GetUpdatesChan(updateConfig)

	updateChan := make(chan tgbotapi.Update)
	go func() {
		for update := range updates {
			updateChan <- update
		}
	}()

	// 启动定时任务
	scheduler.StartScheduler(bot)

	for update := range updateChan {
		go handleUpdate(bot, update)
	}
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if !update.Message.IsCommand() {
		return
	}

	replyMsg := command.HandleCommands(update)

	if _, err := bot.Send(replyMsg); err != nil {
		log.Panic(err)
	}
}
