# Telegram Bot in Go

This document describes the steps to create a simple Telegram bot using Go. The bot will respond with "hello world" to the `/hello` command.

## Prerequisites

*   Go installed on your system.
*   A Telegram account.

## Steps

### 1. Create a new bot on Telegram

1.  Open the Telegram app and search for the "BotFather" user.
2.  Start a chat with BotFather and send the `/newbot` command.
3.  Follow the instructions to choose a name and a username for your bot.
4.  BotFather will provide you with a token. Keep this token safe, as it is required to control your bot.

### 2. Project Setup

1.  Create a new directory for your project.
2.  Inside the project directory, initialize a new Go module:
    ```bash
    go mod init github.com/your-username/telegram-bot
    ```
3.  Install the `go-telegram-bot-api` library:
    ```bash
    go get github.com/go-telegram-bot-api/telegram-bot-api/v5
    ```

### 3. Code

Create a `main.go` file with the following content:

```go
package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "hello":
			msg.Text = "hello world"
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
```

### 4. Running the bot

1.  Set the `TELEGRAM_APITOKEN` environment variable with the token you received from BotFather:
    ```bash
    export TELEGRAM_APITOKEN="YOUR_TELEGRAM_BOT_TOKEN"
    ```
2.  Run the `main.go` file:
    ```bash
    go run main.go
    ```

Your bot should now be running and will respond to the `/hello` command.

### 5. Running with Docker

1.  Make sure you have Docker and Docker Compose installed.
2.  Create a `.env` file in the root of the project with your Telegram Bot Token:
    ```
    TELEGRAM_APITOKEN=YOUR_TELEGRAM_BOT_TOKEN
    ```
3.  Build and run the bot using Docker Compose:
    ```bash
    docker-compose up --build
    ```
