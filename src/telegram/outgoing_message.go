package telegram

type OutgoingMessage interface {
	ChatId() int64
	Text() string
}

type outgoingMessage struct {
	chatId int64
	text   string
}

func (m outgoingMessage) ChatId() int64 {
	return m.chatId
}

func (m outgoingMessage) Text() string {
	return m.text
}

func NewOutgoingMessage(chatId int64, text string) OutgoingMessage {
	return &outgoingMessage{chatId, text}
}
