package telegram

type Handler interface {
	Handle(Update, User, Client) error
}

func NewHandler() Handler {
	return &handler{}
}

type handler struct{}

func (h handler) Handle(update Update, user User, client Client) error {
	message := NewOutgoingMessage(update.Message().ChatId(), update.Message().Text())
	return client.SendMessage(message)
}
