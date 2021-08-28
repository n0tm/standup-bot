package application

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	databaseUrl, telegramAccessToken := "::database url::", "::access token::"
	assert.Equal(t, NewConfig(telegramAccessToken, databaseUrl), &config{telegramAccessToken, databaseUrl})
}

func Test_config_DatabaseDsn(t *testing.T) {
	databaseUrl := "::database url::"
	assert.Equal(t, (&config{databaseDsn: databaseUrl}).DatabaseDsn(), databaseUrl)
}

func Test_config_TelegramBotAccessToken(t *testing.T) {
	telegramAccessToken := "::telegram access token::"
	assert.Equal(t, (&config{telegramBotAccessToken: telegramAccessToken}).TelegramBotAccessToken(), telegramAccessToken)
}
