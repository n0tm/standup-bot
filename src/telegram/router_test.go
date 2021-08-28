package telegram

import (
	"errors"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewRouter(t *testing.T) {
	handlerMock := NewHandlerMock(t)
	authenticator := NewAuthenticatorMock(t)
	assert.Equal(
		t,
		NewRouter(handlerMock, nil, authenticator),
		&router{handlerMock, nil, authenticator},
	)
}

func Test_router_Route_When_Middleware_Fails(t *testing.T) {
	update := NewUpdateMock(t)

	user := NewUserMock(t)

	middleware := NewMiddlewareMock(t)
	expectedError := errors.New("::middleware error::")
	middleware.MiddlewareMock.Expect(update, user).Return(expectedError)

	handler := NewHandlerMock(t)
	client := NewClientMock(t)
	unexpectedError := errors.New("::handler error::")
	handler.HandleMock.Expect(update, user, client).Return(unexpectedError)

	authenticator := NewAuthenticatorMock(t)
	authenticator.AuthenticateMock.Return(user, nil)
	router := &router{handler, []Middleware{middleware}, authenticator}

	assert.Equal(t, router.Route(update, client), expectedError)
}

func Test_router_Route_When_Handler_Fails(t *testing.T) {
	update := NewUpdateMock(t)

	user := NewUserMock(t)

	middleware := NewMiddlewareMock(t)
	middleware.MiddlewareMock.Expect(update, user).Return(nil)

	handler := NewHandlerMock(t)
	client := NewClientMock(t)
	expectedError := errors.New("::handler error::")
	handler.HandleMock.Expect(update, user, client).Return(expectedError)

	authenticator := NewAuthenticatorMock(t)
	authenticator.AuthenticateMock.Return(user, nil)
	router := &router{handler, []Middleware{middleware}, authenticator}

	assert.Equal(t, router.Route(update, client), expectedError)
}

func Test_router_Route_When_Authenticator_Fails(t *testing.T) {
	user := NewUserMock(t)

	update := NewUpdateMock(t)
	middleware := NewMiddlewareMock(t)
	middleware.MiddlewareMock.Expect(update, user).Return(nil)

	handler := NewHandlerMock(t)
	client := NewClientMock(t)
	handler.HandleMock.Expect(update, user, client).Return(nil)

	authenticator := NewAuthenticatorMock(t)
	expectedError := errors.New("::some error::")
	authenticator.AuthenticateMock.Return(nil, expectedError)
	router := &router{handler, []Middleware{middleware}, authenticator}

	assert.Equal(t, router.Route(update, client), expectedError)
}

func Test_router_Route_When_Successfully(t *testing.T) {
	user := NewUserMock(t)

	update := NewUpdateMock(t)
	middleware := NewMiddlewareMock(t)
	middleware.MiddlewareMock.Expect(update, user).Return(nil)

	handler := NewHandlerMock(t)
	client := NewClientMock(t)
	handler.HandleMock.Expect(update, user, client).Return(nil)

	authenticator := NewAuthenticatorMock(t)
	authenticator.AuthenticateMock.Return(user, nil)
	router := &router{handler, []Middleware{middleware}, authenticator}

	assert.Equal(t, router.Route(update, client), nil)
}
