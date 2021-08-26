package telegram

type Update interface {
	Id() int
	Message() OutgoingMessage
}

func NewUpdate(id int, message OutgoingMessage) Update {
	return &update{id, message}
}

type update struct {
	id      int
	message OutgoingMessage
}

func (u update) Id() int {
	return u.id
}

func (u update) Message() OutgoingMessage {
	return u.message
}
