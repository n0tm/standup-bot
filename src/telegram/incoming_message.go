package telegram

type IncomingMessage interface {
	ChatId() int64
	Text() string
	UserId() int
}

type incomingMessage struct {
	userId int
	chatId int64
	text   string
}

func (m incomingMessage) UserId() int {
	return m.userId
}

func (m incomingMessage) ChatId() int64 {
	return m.chatId
}

func (m incomingMessage) Text() string {
	return m.text
}

func NewIncomingMessage(userId int, chatId int64, text string) IncomingMessage {
	return &incomingMessage{userId, chatId, text}
}
