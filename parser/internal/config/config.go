package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"mickaelalliel.com/telebot/parser/internal/config/decoders"
)

type ApplicationConfiguration struct {
	Env                 string              `default:"development" envconfig:"env"`
	ServerPort          int                 `default:"3000" envconfig:"port"`
	BotToken            string              `required:"true" envconfig:"tg_bot_token"`
	WebhookUrl          decoders.URLDecoder `required:"true" envconfig:"tg_webhook_url"`
	AllowedPrivateUsers []string            `default:"[]" envconfig:"allowed_private_users"`
}

type DatabaseConfiguration struct {
	Host     string `default:"localhost" envconfig:"db_host"`
	Port     int    `default:"5432" envconfig:"db_port"`
	Username string `default:"postgres" envconfig:"db_username"`
	Password string `default:"postgrespassword" envconfig:"db_password"`
	Database string `default:"postgres" envconfig:"db_database"`
	TCP      bool   `default:"false" envconfig:"DB_TCP"`
}

var initConfigOnce sync.Once
var AppConfig ApplicationConfiguration = ApplicationConfiguration{}
var DbConfig DatabaseConfiguration = DatabaseConfiguration{}

func ParseConfiguration() {
	ParseConfigurationFromFile(".env")
}

func ParseConfigurationFromFile(envFilePath string) {
	initConfigOnce.Do(func() {
		env := os.Getenv("ENV")
		if env == "development" || env == "" {
			log.Println("Detected development env. Trying to read .env file")
			err := godotenv.Overload(envFilePath)
			if err != nil {
				log.Fatalf("Error reading .env file from path: %s - Error: %s", envFilePath, err.Error())
			}
		}

		err := envconfig.Process("", &AppConfig)
		if err != nil {
			log.Fatalf("Error unmashaling application configuration: %s", err.Error())
		}

		err = envconfig.Process("DB", &DbConfig)
		if err != nil {
			log.Fatalf("Error unmashaling database configuration: %s", err.Error())
		}
	})
}
