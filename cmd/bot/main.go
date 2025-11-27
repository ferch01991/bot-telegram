package main

import (
	"fmt"
	"log"

	"github.com/ferch01991/bot-telegram/internal/bot"
	"github.com/ferch01991/bot-telegram/internal/config"
)

func main() {
	fmt.Println("Bot de Telegram iniciado..............")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	b, err := bot.NewBot(cfg)
	if err != nil {
		fmt.Println("Error al crear el bot de Telegram:")
		log.Panic(err)
	}

	b.Start()
}
