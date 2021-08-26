package telegram

type Listener interface {
	Listen(offset int, Timeout int, limit int) error
}

type listener struct {
	router Router
	client Client
}

func NewListener(router Router, client Client) Listener {
	return &listener{router, client}
}

func (l listener) Listen(offset int, timeout int, limit int) error {
	updates := l.client.GetUpdates(offset, timeout, limit)
	for update := range updates {
		if err := l.router.Route(update, l.client); err != nil {
			return err
		}
	}

	return nil
}
