package initializers

import (
	"fmt"
	"log"
	"reflect"

	"github.com/spf13/viper"
)

type Environment string

const (
	DevelopmentEnv Environment = "development"
	ProductionEnv  Environment = "production"
)

type Config struct {
	PORT                 string      `mapstructure:"PORT"`
	DB_URL               string      `mapstructure:"DB_URL"`
	JWT_SECRET           string      `mapstructure:"JWT_SECRET"`
	ENV                  Environment `mapstructure:"ENV"`
	FRONTEND_URL         string      `mapstructure:"FRONTEND_URL"`
	BACKEND_URL          string      `mapstructure:"BACKEND_URL"`
	SENDGRID_KEY         string      `mapstructure:"SENDGRID_KEY"`
	GOOGLE_CLIENT_ID     string      `mapstructure:"GOOGLE_CLIENT_ID"`
	GOOGLE_CLIENT_SECRET string      `mapstructure:"GOOGLE_CLIENT_SECRET"`
	GOOGLE_OAUTH_STATE   string      `mapstructure:"GOOGLE_OAUTH_STATE"`
}

var CONFIG Config

func LoadEnv() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&CONFIG)
	if err != nil {
		log.Fatal(err)
	}

	requiredKeys := getRequiredKeys(CONFIG)
	missingKeys := checkMissingKeys(requiredKeys, CONFIG)

	if len(missingKeys) > 0 {
		err := fmt.Errorf("following environment variables not found: %v", missingKeys)
		log.Fatal(err)
	}

	if CONFIG.ENV != DevelopmentEnv && CONFIG.ENV != ProductionEnv {
		err := fmt.Errorf("invalid ENV value: %s", CONFIG.ENV)
		log.Fatal(err)
	}
}

func getRequiredKeys(config Config) []string {
	requiredKeys := []string{}
	configType := reflect.TypeOf(config)

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		tag := field.Tag.Get("mapstructure")
		if tag != "" {
			requiredKeys = append(requiredKeys, tag)
		}
	}

	return requiredKeys
}

func checkMissingKeys(requiredKeys []string, config Config) []string {
	missingKeys := []string{}

	configValue := reflect.ValueOf(config)
	for _, key := range requiredKeys {
		value := configValue.FieldByName(key).String()
		if value == "" {
			missingKeys = append(missingKeys, key)
		}
	}

	return missingKeys
}
