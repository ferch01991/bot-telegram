package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Bot de Telegram iniciado..............")

	if err := godotenv.Load(); err != nil {
		log.Println(".env no encontrado; usando variables de entorno del sistema")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		fmt.Println("Error al crear el bot de Telegram:")
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("Received update: %+v", update) // Debug log for all updates
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		log.Printf("Received message from %s: %s", update.Message.From.UserName, update.Message.Text) // Debug log for messages

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		log.Printf("Received command: %s", update.Message.Command()) // Debug log for commands

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "hello":
			msg.Text = "hello world!"
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
