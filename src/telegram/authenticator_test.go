package telegram

import (
	"errors"
	"github.com/gojuno/minimock/v3"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewAuthenticator(t *testing.T) {
	userRepository := NewUserRepositoryMock(t)
	assert.Equal(t, NewAuthenticator(userRepository), &authenticator{userRepository})
}

func Test_authenticator_Authenticate_When_Fails(t *testing.T) {
	mc := minimock.NewController(t)

	userRepository := NewUserRepositoryMock(mc)
	authenticator := authenticator{userRepository}

	userId := 0
	message := NewIncomingMessageMock(mc)
	message.UserIdMock.Return(userId)

	update := NewUpdateMock(mc)
	update.MessageMock.Return(message)

	user, err := authenticator.Authenticate(update)
	assert.Equal(t, user, nil)
	assert.Equal(t, err, errors.New("can't identificate user"))
}

func Test_authenticator_Authenticate_When_Success(t *testing.T) {
	mc := minimock.NewController(t)

	userId := 12345
	userMock := NewUserMock(mc)
	userMock.GetIdMock.Return(uint(userId))

	userRepository := NewUserRepositoryMock(mc)
	userRepository.GetByIdMock.Expect(uint(userId)).Return(userMock)

	authenticator := authenticator{userRepository}

	message := NewIncomingMessageMock(mc)
	message.UserIdMock.Return(userId)

	update := NewUpdateMock(mc)
	update.MessageMock.Return(message)

	user, err := authenticator.Authenticate(update)
	assert.Equal(t, user, userMock)
	assert.Equal(t, err, nil)
}
