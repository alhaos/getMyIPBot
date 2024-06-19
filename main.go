package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"io"
	"log"
	"net/http"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("6270193035:AAGoGyi6SHiY9qldyglHVS063F7qn8VG6k8")
	if err != nil {
		panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message.Text == "GetMyIP" { // If we got a message

			response, err := http.Get("https://icanhazip.com/")
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "error")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}

			body, err := io.ReadAll(response.Body)
			_ = response.Body.Close()

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(body))
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}

}
