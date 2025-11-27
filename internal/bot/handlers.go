package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// HandleHello responds to the "hello" command
func HandleHello(update tgbotapi.Update) string {
	return "hello world, perrito miau miau!"
}

// HandleDefault responds to unknown commands
func HandleDefault(update tgbotapi.Update) string {
	return "I don't know that command"
}
