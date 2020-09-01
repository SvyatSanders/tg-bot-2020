package main

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	webHook = "https://sandersbot.herokuapp.com/"
)

func main() {
	port := os.Getenv("PORT")

	go func() {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()

	bot, err := tgbotapi.NewBotAPI(botConfig.token)
	if err != nil {
		log.Fatal("creating bot fail: ", err)
	}
	log.Println("bot has been born ;)")

	if _, err = bot.SetWebhook(tgbotapi.NewWebhook(webHook)); err != nil {
		log.Fatalf("fail during the webhook process: %v, error: %v", webHook, err)
	}
	log.Println("webhook created")

	updates := bot.ListenForWebhook("/")
	for update := range updates {
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)); err != nil {
			log.Print(err)
		}
	}

}
