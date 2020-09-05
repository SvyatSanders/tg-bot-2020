package main

import (
	"fmt"
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

	// создаем botApi
	bot, err := tgbotapi.NewBotAPI(botConfig.tgToken)
	if err != nil {
		log.Fatal("creating bot fail: ", err)
	}
	log.Println("bot has been born ;)")

	// создаем webhook - механизм получения уведомлений об определённых событиях
	if _, err = bot.SetWebhook(tgbotapi.NewWebhook(webHook)); err != nil {
		log.Fatalf("fail during the webhook process: %v, error: %v", webHook, err)
	}
	log.Println("webhook created")

	// хэндлер для прослушивания событий, переданных через webhook
	updates := bot.ListenForWebhook("/")
	command := "vk"
	groupID := "-15365973"

	for update := range updates {
		text := fmt.Sprintf("command is not %v", command)
		if command != update.Message.Text {
			if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text)); err != nil {
				log.Printf("request failed: %v", err)
			}
			continue
		}

		items, err := getPostsQuery(groupID, botConfig.vkServiceKey)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, item := range items {
			if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, item.Text)); err != nil {
				log.Printf("request failed: %v", err)
			}
		}
	}
}
