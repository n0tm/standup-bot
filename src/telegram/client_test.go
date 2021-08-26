package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	api := &tgbotapi.BotAPI{}
	assert.Equal(t, NewClient(api), &client{api})
}

func Test_client_GetUpdates(t *testing.T) {
	t.Skip("Пока не понятно как тестировать зависимости библиотек без интерфейсов")
}

func Test_client_SendMessage(t *testing.T) {
	t.Skip("Пока не понятно как тестировать зависимости библиотек без интерфейсов")
}
