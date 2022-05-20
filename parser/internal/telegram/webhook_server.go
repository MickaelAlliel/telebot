package telegram

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"mickaelalliel.com/telebot/parser/internal/config"
)

func NewTelegramWebhookServerOrFail() tgbotapi.UpdatesChannel {
	bot, err := tgbotapi.NewBotAPI(config.AppConfig.BotToken)
	if err != nil {
		log.Fatal(err.Error())
	}

	if config.AppConfig.Env == "development" {
		bot.Debug = true
	}

	log.Printf("Starting up bot: %s", bot.Self.UserName)
	webhookUrl := url.URL(config.AppConfig.WebhookUrl)
	webhookUrl.Path = path.Join(webhookUrl.Path, config.AppConfig.BotToken)
	wh, _ := tgbotapi.NewWebhook(webhookUrl.String())

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.AppConfig.ServerPort), nil)

	return updates
}
