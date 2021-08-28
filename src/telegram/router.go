package telegram

type Router interface {
	Route(Update, Client) error
}

type router struct {
	handler       Handler
	middlewares   []Middleware
	authenticator Authenticator
}

func NewRouter(handler Handler, middlewares []Middleware, authenticator Authenticator) Router {
	return &router{handler, middlewares, authenticator}
}

func (r router) Route(update Update, client Client) error {
	user, err := r.authenticator.Authenticate(update)
	if err != nil {
		return err
	}

	for _, middleware := range r.middlewares {
		if err := middleware.Middleware(update, user); err != nil {
			return err
		}
	}

	if err := r.handler.Handle(update, user, client); err != nil {
		return err
	}

	return nil
}
