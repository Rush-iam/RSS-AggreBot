package tg_client

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/ratelimit"
	"log"
)

func New(token string) *tgbotapi.BotAPI {
	tgClient, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Can't get TG client: %v", err)
	} else {
		log.Printf("Connected to Telegram API as @%s", tgClient.Self.UserName)
	}
	return tgClient
}

var rl = ratelimit.New(30)

func SendMessage(tgClient *tgbotapi.BotAPI, userId int64, text string, markup *tgbotapi.InlineKeyboardMarkup) error {
	if text == "" {
		return nil
	}
	rl.Take()
	reply := tgbotapi.NewMessage(userId, text)
	reply.ReplyMarkup = markup
	_, err := tgClient.Send(reply)
	return err
}

func SendCallbackAnswer(tgClient *tgbotapi.BotAPI, queryId, text string) error {
	answer := tgbotapi.NewCallback(queryId, text)
	_, err := tgClient.Send(answer)
	return err
}

func UpdateMessage(tgClient *tgbotapi.BotAPI, userId int64, messageId int, text string, markup *tgbotapi.InlineKeyboardMarkup) error {
	if text == "" {
		return nil
	}

	var reply tgbotapi.EditMessageTextConfig
	if markup != nil {
		reply = tgbotapi.NewEditMessageTextAndMarkup(userId, messageId, text, *markup)
	} else {
		reply = tgbotapi.NewEditMessageText(userId, messageId, text)
	}
	_, err := tgClient.Send(reply)
	return err
}
