package telegram

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewAuthorizationMiddleware(t *testing.T) {
	expected, actual := &authorizationMiddleware{}, NewAuthorizationMiddleware()
	assert.Equal(t, actual, expected)
}

func Test_authorizationMiddleware_Middleware(t *testing.T) {
	middleware := &authorizationMiddleware{}
	assert.Equal(t, middleware.Middleware(NewUpdateMock(t)), nil)
}
