package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const updatesChannelBufferSize int = 100

type Client interface {
	GetUpdates(offset, timeout, limit int) chan Update
	SendMessage(OutgoingMessage) error
}

func NewClient(api *tgbotapi.BotAPI) Client {
	return &client{api}
}

type client struct {
	api *tgbotapi.BotAPI
}

func (c client) SendMessage(m OutgoingMessage) error {
	_, err := c.api.Send(tgbotapi.NewMessage(m.ChatId(), m.Text()))
	return err
}

func (c client) GetUpdates(offset, timeout, limit int) chan Update {
	channel := make(chan Update, updatesChannelBufferSize)

	go func() {
		updateConfig := tgbotapi.NewUpdate(offset)
		updateConfig.Timeout = timeout
		updateConfig.Limit = limit

		updates, _ := c.api.GetUpdatesChan(updateConfig)
		for update := range updates {
			channel <- NewUpdate(update.UpdateID, NewIncomingMessage(update.Message.From.ID, update.Message.Chat.ID, update.Message.Text))
		}
	}()

	return channel
}
