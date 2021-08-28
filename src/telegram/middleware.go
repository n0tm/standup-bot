package telegram

type Middleware interface {
	Middleware(Update, User) error
}
