package telegram

import (
	"context"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/exp/slices"
	"mickaelalliel.com/telebot/parser/internal/config"
	"mickaelalliel.com/telebot/parser/internal/ent"
	"mickaelalliel.com/telebot/parser/internal/parsers"
)

func HandleTelegramUpdates(ctx context.Context, updateChannel tgbotapi.UpdatesChannel) {
	bot, ok := ctx.Value(config.ContextKeyBot).(*tgbotapi.BotAPI)
	if !ok {
		log.Fatal("failed to get bot api client from context")
		return
	}
	db, ok := ctx.Value(config.ContextKeyDbClient).(*ent.Client)
	if !ok {
		log.Fatal("failed to get db client from context")
		return
	}

	for update := range updateChannel {
		go HandleTelegramUpdate(ctx, bot, db, &update)
	}
}

func replyToMessage(bot *tgbotapi.BotAPI, upd *tgbotapi.Update, message string) (tgbotapi.Message, error) {
	reply := tgbotapi.NewMessage(upd.FromChat().ID, message)
	msg, error := bot.Send(reply)
	return msg, error
}

func HandleTelegramUpdate(ctx context.Context, bot *tgbotapi.BotAPI, db *ent.Client, upd *tgbotapi.Update) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
			replyToMessage(bot, upd, "fatal error! check logs.")
		}
	}()

	if upd.FromChat().IsPrivate() {
		if !slices.Contains(config.AppConfig.AllowedPrivateUsers, upd.Message.From.UserName) {
			replyToMessage(bot, upd, config.ErrorMessage_UnauthorizedUser)
			return
		}
	}

	if upd.Message.Command() == "start" {
		replyToMessage(bot, upd, config.StartMessage)
		return
	}

	expense, err := parsers.ParseExpenseMessage(upd.Message)
	if err != nil {
		log.Println(err.Error())
		replyToMessage(bot, upd, config.ErrorMessage_BadRequest)
		return
	}

	saveCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	_, err = expense.SaveEnt(saveCtx, db)
	if err != nil {
		log.Println(err.Error())
		replyToMessage(bot, upd, config.ErrorMessage_FailedToSave)
		return
	}
	replyToMessage(bot, upd, config.SuccessMessage_Saved)
}
