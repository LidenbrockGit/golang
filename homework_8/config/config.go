package config

import (
	"errors"
	"flag"
	"github.com/kelseyhightower/envconfig"
	"net/url"
)

type Config struct {
	Port        uint64 `ignored:"true"`
	DbUrl       string `required:"true" envconfig:"DB_URL"`
	JaegerUrl   string `required:"true" envconfig:"JAEGER_URL"`
	SentryUrl   string `required:"true" envconfig:"SENTRY_URL"`
	KafkaBroker string `default:"kafka:9092" envconfig:"KAFKA_BROKER"`
	SomeAppId   string `ignored:"true"`
	SomeAppKey  string `ignored:"true"`
}

type ConfUrl struct {
	url string
}

func (u ConfUrl) decode() (string, error) {
	var response string
	parseResult, err := url.Parse(u.url)
	if parseResult != nil {
		response = parseResult.String()
	}
	return response, err
}

func Read() (Config, error) {
	configData := Config{}

	err := readFlags(&configData)
	if err != nil {
		return Config{}, err
	}
	err = readEnv(&configData)
	if err != nil {
		return Config{}, err
	}

	return configData, nil
}

// Чтение флагов
func readFlags(confData *Config) error {
	fPort := flag.Uint64("port", 8080, "Port number")
	fSomeAppId := flag.String("some_app_id", "", "Some App ID")
	fSomeAppKey := flag.String("some_app_key", "", "Some App Key")

	flag.Parse()

	if *fPort == 0 {
		return errors.New("port can't be empty")
	}
	if *fSomeAppId == "" {
		return errors.New("some_app_id can't be empty")
	}
	if *fSomeAppKey == "" {
		return errors.New("some_app_key can't be empty")
	}

	confData.Port = *fPort
	confData.SomeAppId = *fSomeAppId
	confData.SomeAppKey = *fSomeAppKey

	return nil
}

// Чтение переменных окружения
func readEnv(confData *Config) error {
	err := envconfig.Process("", confData)
	if err != nil {
		return err
	}

	_, err = ConfUrl{url: confData.DbUrl}.decode()
	if err != nil {
		return errors.New("DB_URL is incorrect")
	}

	_, err = ConfUrl{url: confData.JaegerUrl}.decode()
	if err != nil {
		return errors.New("JAEGER_URL is incorrect")
	}

	_, err = ConfUrl{url: confData.SentryUrl}.decode()
	if err != nil {
		return errors.New("SENTRY_URL is incorrect")
	}

	return nil
}
