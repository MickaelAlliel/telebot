package main

import (
	"context"

	"mickaelalliel.com/telebot/parser/internal/config"
	"mickaelalliel.com/telebot/parser/internal/db"
	"mickaelalliel.com/telebot/parser/internal/telegram"
)

func main() {
	config.ParseConfiguration()
	dbClient := db.NewDatabaseOrFail()
	defer dbClient.Close()

	updates, bot := telegram.NewTelegramWebhookServerOrFail()
	ctx := context.WithValue(context.Background(), config.ContextKeyBot, bot)
	ctx = context.WithValue(ctx, config.ContextKeyDbClient, dbClient)
	telegram.HandleTelegramUpdates(ctx, updates)
}
