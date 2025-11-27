package bot

import (
	"log"

	"github.com/ferch01991/bot-telegram/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Bot represents the Telegram bot
type Bot struct {
	api *tgbotapi.BotAPI
	cfg *config.Config
}

// NewBot creates a new Bot instance
func NewBot(cfg *config.Config) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.TelegramAPIToken)
	if err != nil {
		return nil, err
	}

	api.Debug = cfg.Debug
	log.Printf("Authorized on account %s", api.Self.UserName)

	return &Bot{
		api: api,
		cfg: cfg,
	}, nil
}

// Start starts the bot and listens for updates
func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("Received update: %+v", update)

		if update.Message == nil {
			continue
		}

		log.Printf("Received message from %s: %s", update.Message.From.UserName, update.Message.Text)

		if !update.Message.IsCommand() {
			continue
		}

		log.Printf("Received command: %s", update.Message.Command())

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "hello":
			msg.Text = HandleHello(update)
		default:
			msg.Text = HandleDefault(update)
		}

		if _, err := b.api.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
