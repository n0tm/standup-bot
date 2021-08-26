package telegram

import (
	"errors"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewRouter(t *testing.T) {
	handlerMock := NewHandlerMock(t)
	middlewareMocks := []Middleware{NewMiddlewareMock(t)}
	assert.Equal(t, NewRouter(handlerMock, middlewareMocks), &router{handlerMock, middlewareMocks})
}

func Test_router_Route_When_Middleware_Fails(t *testing.T) {
	update := NewUpdateMock(t)

	middleware := NewMiddlewareMock(t)
	expectedError := errors.New("::middleware error::")
	middleware.MiddlewareMock.Expect(update).Return(expectedError)

	handler := NewHandlerMock(t)
	client := NewClientMock(t)
	unexpectedError := errors.New("::handler error::")
	handler.HandleMock.Expect(update, client).Return(unexpectedError)

	router := &router{handler, []Middleware{middleware}}

	assert.Equal(t, router.Route(update, client), expectedError)
}

func Test_router_Route_When_Handler_Fails(t *testing.T) {
	update := NewUpdateMock(t)

	middleware := NewMiddlewareMock(t)
	middleware.MiddlewareMock.Expect(update).Return(nil)

	handler := NewHandlerMock(t)
	client := NewClientMock(t)
	expectedError := errors.New("::handler error::")
	handler.HandleMock.Expect(update, client).Return(expectedError)

	router := &router{handler, []Middleware{middleware}}

	assert.Equal(t, router.Route(update, client), expectedError)
}

func Test_router_Route_When_Middleware_Successfully(t *testing.T) {
	update := NewUpdateMock(t)

	middleware := NewMiddlewareMock(t)
	middleware.MiddlewareMock.Expect(update).Return(nil)

	handler := NewHandlerMock(t)
	client := NewClientMock(t)
	handler.HandleMock.Expect(update, client).Return(nil)

	router := &router{handler, []Middleware{middleware}}

	assert.Equal(t, router.Route(update, client), nil)
}
