package bot

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TestHandleHello(t *testing.T) {
	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: "/hello",
			From: &tgbotapi.User{
				UserName: "testuser",
			},
		},
	}

	expected := "hello world, perrito miau miau!"
	result := HandleHello(update)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestHandleDefault(t *testing.T) {
	update := tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: "/unknown",
		},
	}

	expected := "I don't know that command"
	result := HandleDefault(update)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
