package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"mickaelalliel.com/telebot/parser/internal/config/decoders"
)

type Config struct {
	Env        string              `default:"development" envconfig:"env"`
	ServerPort int                 `default:"3000" envconfig:"port"`
	BotToken   string              `required:"true" envconfig:"tg_bot_token"`
	WebhookUrl decoders.URLDecoder `required:"true" envconfig:"tg_webhook_url"`
}

var initConfigOnce sync.Once
var AppConfig Config = Config{}

func ParseConfiguration() {
	ParseConfigurationFromFile(".env")
}

func ParseConfigurationFromFile(envFilePath string) {
	initConfigOnce.Do(func() {
		env := os.Getenv("ENV")
		if env == "development" {
			log.Println("Detected development env. Trying to read .env file")
			err := godotenv.Overload(envFilePath)
			if err != nil {
				log.Fatalf("Error reading .env file from path: %s - Error: %s", envFilePath, err.Error())
			}
		}

		err := envconfig.Process("", &AppConfig)
		if err != nil {
			log.Fatalf("Error unmashaling configuration: %s", err.Error())
		}
	})
}
