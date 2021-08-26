package telegram

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewIncomingMessage(t *testing.T) {
	var chatId int64 = 1234
	text := "::some text::"

	expected, actual := &incomingMessage{chatId, text}, NewIncomingMessage(chatId, text)

	assert.Equal(t, actual, expected)
}

func Test_incomingMessage_GetChatId(t *testing.T) {
	var chatId int64 = 1234

	incomingMessage := &incomingMessage{chatId: chatId}
	assert.Equal(t, incomingMessage.ChatId(), chatId)
}

func Test_incomingMessage_GetText(t *testing.T) {
	text := "::some text::"

	incomingMessage := &incomingMessage{text: text}
	assert.Equal(t, incomingMessage.Text(), text)
}
