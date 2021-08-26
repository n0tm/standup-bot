package telegram

type Middleware interface {
	Middleware(Update) error
}

func NewAuthorizationMiddleware() Middleware {
	return &authorizationMiddleware{}
}

type authorizationMiddleware struct {
}

func (m authorizationMiddleware) Middleware(u Update) error {
	return nil
}
