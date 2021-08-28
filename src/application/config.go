package application

type Config interface {
	TelegramBotAccessToken() string
	DatabaseDsn() string
}

func NewConfig(telegramBotAccessToken, databaseDsn string) Config {
	return &config{telegramBotAccessToken, databaseDsn}
}

type config struct {
	telegramBotAccessToken string
	databaseDsn            string
}

func (c config) TelegramBotAccessToken() string {
	return c.telegramBotAccessToken
}

func (c config) DatabaseDsn() string {
	return c.databaseDsn
}
