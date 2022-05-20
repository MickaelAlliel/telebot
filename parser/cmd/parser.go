package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"mickaelalliel.com/telebot/parser/internal/config"
	"mickaelalliel.com/telebot/parser/internal/db"
)

func main() {
	config.ParseConfigurationFromFile("../.env")
	client := db.NewDatabaseOrFail()
	defer client.Close()

	bot, err := tgbotapi.NewBotAPI(config.AppConfig.BotToken)
	if err != nil {
		log.Fatal(err.Error())
	}

	bot.Debug = true

	log.Printf("Starting up bot: %s", bot.Self.UserName)

	wh, _ := tgbotapi.NewWebhook(path.Join(config.AppConfig.WebhookUrl.Path, config.AppConfig.BotToken))

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

	for update := range updates {
		log.Printf("%+v\n", update.Message.Text)
	}
}
