package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	const url = "https://icanhazip.com"

	token := os.Getenv("IP_BOT_TOKEN")
	if token == "" {
		log.Fatalf("env variable IP_BOT_TOKEN not found")
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("unbale to get updates for chanal: %s", err)
	}

	log.Println("Application started")

	for update := range updates {
		if update.Message.Text == "GetMyIP" {
			response, err2 := http.Get(url)
			if err2 != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "error: "+err2.Error())
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				continue
			}

			body, err3 := io.ReadAll(response.Body)
			if err3 != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "unable to read msg body")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				continue
			}

			_ = response.Body.Close()

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(body))
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)

			log.Println("request precessed from user:", update.Message.From.UserName)
		}
	}
}
