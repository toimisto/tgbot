package main

import (
	"log"
	"strings"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

    token, exists := os.LookupEnv("BOT_TOKEN")

    if !exists {
		log.Panic("Token for bot not found.")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		message := update.Message.Text

		if strings.Contains(message, "toi") || strings.Contains(message, "Toi") {
			response := strings.Replace(message, "toi", "töi", -1)
			response = strings.Replace(response, "Toi", "Töi", -1)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Tarkoitit kai: " + response)
			bot.Send(msg)
		}

	}
}
