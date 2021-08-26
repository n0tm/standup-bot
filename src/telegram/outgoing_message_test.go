package telegram

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewOutgoingMessage(t *testing.T) {
	var chatId int64 = 123
	text := "::some text::"

	expected, actual := &outgoingMessage{chatId, text}, NewOutgoingMessage(chatId, text)

	assert.Equal(t, expected, actual)
}

func Test_outgoingMessage_GetChatId(t *testing.T) {
	var chatId int64 = 123

	message := &outgoingMessage{chatId: chatId}

	assert.Equal(t, message.ChatId(), chatId)
}

func Test_outgoingMessage_GetText(t *testing.T) {
	text := "::some text::"

	message := &outgoingMessage{text: text}

	assert.Equal(t, message.Text(), text)
}
