package telegram

type Router interface {
	Route(Update, Client) error
}

type router struct {
	handler     Handler
	middlewares []Middleware
}

func NewRouter(handler Handler, middlewares []Middleware) Router {
	return &router{handler, middlewares}
}

func (r router) Route(update Update, client Client) error {
	for _, middleware := range r.middlewares {
		if err := middleware.Middleware(update); err != nil {
			return err
		}
	}

	if err := r.handler.Handle(update, client); err != nil {
		return err
	}

	return nil
}
