package main

import (
	"log"

	"mickaelalliel.com/telebot/parser/internal/config"
	"mickaelalliel.com/telebot/parser/internal/db"
	"mickaelalliel.com/telebot/parser/internal/telegram"
)

func main() {
	config.ParseConfigurationFromFile("../.env")
	client := db.NewDatabaseOrFail()
	defer client.Close()

	updates := telegram.NewTelegramWebhookServerOrFail()
	for update := range updates {
		log.Printf("%+v\n", update.Message.Text)
	}
}
