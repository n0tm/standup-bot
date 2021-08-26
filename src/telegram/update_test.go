package telegram

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewUpdate(t *testing.T) {
	id := 123
	message := NewIncomingMessageMock(t)

	expected, actual := &update{id, message}, NewUpdate(id, message)

	assert.Equal(t, actual, expected)
}

func Test_update_GetId(t *testing.T) {
	expectedId := 123
	update := &update{id: expectedId}
	assert.Equal(t, update.Id(), expectedId)
}

func Test_update_GetMessage(t *testing.T) {
	expectedMessage := NewIncomingMessageMock(t)
	update := &update{message: expectedMessage}
	assert.Equal(t, update.Message(), expectedMessage)
}
