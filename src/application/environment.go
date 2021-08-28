package application

import (
	"github.com/joho/godotenv"
	"os"
)

type Environment interface {
	Load(path string) error
	TelegramBotAccessToken() string
	DatabaseDsn() string
}

const (
	KeyNameTelegramBotAccessToken = "TELEGRAM_BOT_ACCESS_TOKEN"
	KeyNameDatabaseDsn            = "DATABASE_DSN"
)

func NewEnvironment() Environment {
	return &environment{}
}

type environment struct{}

func (e environment) Load(path string) error {
	if err := godotenv.Load(path); err != nil {
		return err
	}

	return nil
}

func (e environment) TelegramBotAccessToken() string {
	return getEnv(KeyNameTelegramBotAccessToken, "")
}

func (e environment) DatabaseDsn() string {
	return getEnv(KeyNameDatabaseDsn, "")
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
