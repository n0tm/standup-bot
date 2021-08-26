package telegram

type Handler interface {
	Handle(Update, Client) error
}

func NewHandler() Handler {
	return &handler{}
}

type handler struct{}

func (h handler) Handle(update Update, client Client) error {
	message := NewOutgoingMessage(update.Message().ChatId(), update.Message().Text())
	return client.SendMessage(message)
}
