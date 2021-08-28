package telegram

import (
	"github.com/gojuno/minimock/v3"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewHandler(t *testing.T) {
	expected, actual := &handler{}, NewHandler()
	assert.Equal(t, actual, expected)
}

func Test_handler_Handle(t *testing.T) {
	mc := minimock.NewController(t)

	message := NewIncomingMessageMock(mc)

	var chatId int64 = 123
	message.ChatIdMock.Return(chatId)

	text := "::some text::"
	message.TextMock.Return(text)

	update := NewUpdateMock(mc)
	update.MessageMock.Return(message)

	updateId := 1
	update.IdMock.Return(updateId)

	client := NewClientMock(mc)
	client.SendMessageMock.Expect(&outgoingMessage{chatId, text}).Return(nil)

	handler := &handler{}
	assert.Equal(t, handler.Handle(update, NewUserMock(mc), client), nil)
}
