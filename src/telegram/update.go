package telegram

type Update interface {
	Id() int
	Message() IncomingMessage
}

func NewUpdate(id int, message IncomingMessage) Update {
	return &update{id, message}
}

type update struct {
	id      int
	message IncomingMessage
}

func (u update) Id() int {
	return u.id
}

func (u update) Message() IncomingMessage {
	return u.message
}
