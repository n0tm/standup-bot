package telegram

import (
	"errors"
	"github.com/gojuno/minimock/v3"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewListener(t *testing.T) {
	router := NewRouterMock(t)
	client := NewClientMock(t)

	expected, actual := &listener{router, client}, NewListener(router, client)

	assert.Equal(t, expected, actual, "Not equal listeners")
}

func Test_listener_Listen_When_Successfully(t *testing.T) {
	mc := minimock.NewController(t)

	client := NewClientMock(mc)

	router := NewRouterMock(mc)
	update := NewUpdateMock(mc)
	router.RouteMock.Expect(update, client).Return(nil)

	offset := 0
	timeout := 0
	limit := 0
	updatesChannel := make(chan Update)
	client.GetUpdatesMock.Expect(offset, timeout, limit).Return(updatesChannel)

	listener := &listener{router, client}

	go func() {
		updatesChannel <- update
		defer close(updatesChannel)
	}()

	assert.Equal(t, listener.Listen(offset, timeout, limit), nil)
}

func Test_listener_Listen_When_Fails(t *testing.T) {
	mc := minimock.NewController(t)

	client := NewClientMock(mc)

	router := NewRouterMock(mc)
	update := NewUpdateMock(mc)
	expectedError := errors.New("::some error::")
	router.RouteMock.Expect(update, client).Return(expectedError)

	offset := 0
	timeout := 0
	limit := 0
	updatesChannel := make(chan Update)
	client.GetUpdatesMock.Expect(offset, timeout, limit).Return(updatesChannel)

	listener := &listener{router, client}

	go func() {
		updatesChannel <- update
		defer close(updatesChannel)
	}()

	assert.Equal(t, listener.Listen(offset, timeout, limit), expectedError)
}
