package telegram

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewIncomingMessage(t *testing.T) {
	var chatId int64 = 1234
	var userId = 12345
	text := "::some text::"

	expected, actual := &incomingMessage{userId, chatId, text}, NewIncomingMessage(userId, chatId, text)

	assert.Equal(t, actual, expected)
}

func Test_incomingMessage_ChatId(t *testing.T) {
	var chatId int64 = 1234

	incomingMessage := &incomingMessage{chatId: chatId}
	assert.Equal(t, incomingMessage.ChatId(), chatId)
}

func Test_incomingMessage_Text(t *testing.T) {
	text := "::some text::"

	incomingMessage := &incomingMessage{text: text}
	assert.Equal(t, incomingMessage.Text(), text)
}

func Test_incomingMessage_UserId(t *testing.T) {
	userId := 12345

	incomingMessage := &incomingMessage{userId: userId}
	assert.Equal(t, incomingMessage.UserId(), userId)
}
