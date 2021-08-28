package main

import (
	"github.com/apex/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.uber.org/dig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"standup-bot/src/application"
	"standup-bot/src/telegram"
)

func main() {
	app := application.Instance()

	container, err := loadContainer()

	if err != nil {
		log.Errorf("error during application loading", err)
	}

	app.SetContainer(container)

	err = run(app)
	if err != nil {
		log.Errorf("error during running", err)
	}
}

func run(app application.Application) error {
	return app.Container().Invoke(func(listener telegram.Listener) error {
		if err := listener.Listen(0, 60, 0); err != nil {
			return err
		}

		return nil
	})
}

func loadContainer() (application.Container, error) {
	container := application.NewContainer(dig.New())

	container, err := loadApplication(container)
	if err != nil {
		return nil, err
	}

	container, err = loadTelegram(container)
	if err != nil {
		return nil, err
	}

	return container, nil
}

func loadApplication(container application.Container) (application.Container, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=standup-bot password=standup-bot dbname=standup-bot port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&telegram.UserModel{},
	)

	container.Provide(func() application.Storage {
		return application.NewStorage(db)
	})

	return container, nil
}

func loadTelegram(container application.Container) (application.Container, error) {
	if err := container.Provide(telegram.NewListener); err != nil {
		return nil, err
	}

	if err := container.Provide(telegram.NewHandler); err != nil {
		return nil, err
	}

	if err := container.Provide(telegram.NewAuthenticator); err != nil {
		return nil, err
	}

	if err := container.Provide(telegram.NewUserRepository); err != nil {
		return nil, err
	}

	err := container.Provide(func() telegram.Client {
		api, _ := tgbotapi.NewBotAPI("1939189655:AAHOvN8_3MqJA1yGTmh5Z5lMIQ6c_TypqOM")
		return telegram.NewClient(api)
	})

	if err != nil {
		return nil, err
	}

	err = container.Provide(func(handler telegram.Handler, authenticator telegram.Authenticator) telegram.Router {
		return telegram.NewRouter(handler, nil, authenticator)
	})

	if err != nil {
		return nil, err
	}

	return container, nil
}
