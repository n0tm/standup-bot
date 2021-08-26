package main

import (
	"github.com/apex/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/dig"
	"standup-bot/src/application"
	"standup-bot/src/telegram"
)

func main() {
	app := application.Instance()

	container := application.NewContainer(dig.New())
	container = loadContainer(container)

	app.SetContainer(container)

	app.Container().Invoke(func(listener telegram.Listener) {
		if err := listener.Listen(0, 60, 0); err != nil {
			log.Errorf("error during listening", err)
		}
	})
}

func loadContainer(container application.Container) application.Container {
	container.Provide(telegram.NewListener)
	container.Provide(telegram.NewHandler)
	container.Provide(telegram.NewAuthorizationMiddleware)

	container.Provide(func() telegram.Client {
		api, _ := tgbotapi.NewBotAPI("1939189655:AAHOvN8_3MqJA1yGTmh5Z5lMIQ6c_TypqOM")
		return telegram.NewClient(api)
	})

	container.Provide(func(handler telegram.Handler) telegram.Router {
		return telegram.NewRouter(handler, []telegram.Middleware{telegram.NewAuthorizationMiddleware()})
	})

	return container
}
