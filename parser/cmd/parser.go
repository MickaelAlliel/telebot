package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Overload("../.env")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}

	serverPort, exists := os.LookupEnv("PORT")
	if !exists {
		log.Println("Using default port :3000")
		serverPort = "3000"
	}

	tgBotToken, exists := os.LookupEnv("TG_BOT_TOKEN")
	if !exists {
		log.Fatal("Missing TG_BOT_TOKEN environment variable")
	}

	tgWebhookUrl, exists := os.LookupEnv("TG_WEBHOOK_URL")
	if !exists {
		log.Fatal("Missing TG_WEBHOOK_URL environment variable")
	}

	bot, err := tgbotapi.NewBotAPI(tgBotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Starting up bot: %s", bot.Self.UserName)

	whUrl, err := url.Parse(tgWebhookUrl)
	if err != nil {
		log.Fatal("Failed to parse webhook url")
	}
	whUrl.Path = path.Join(whUrl.Path, bot.Token)
	wh, _ := tgbotapi.NewWebhook(whUrl.String())

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
	go http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", serverPort), nil)

	for update := range updates {
		log.Printf("%+v\n", update.Message.Text)
	}
}
