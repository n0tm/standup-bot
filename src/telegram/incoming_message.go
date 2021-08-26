package telegram

type IncomingMessage interface {
	ChatId() int64
	Text() string
}

type incomingMessage struct {
	chatId int64
	text   string
}

func (m incomingMessage) ChatId() int64 {
	return m.chatId
}

func (m incomingMessage) Text() string {
	return m.text
}

func NewIncomingMessage(chatId int64, text string) IncomingMessage {
	return &incomingMessage{chatId, text}
}
