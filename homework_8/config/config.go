package config

import (
	"errors"
	"flag"
	"github.com/kelseyhightower/envconfig"
	"net/url"
)

type Config struct {
	Port        uint64
	DbUrl       string
	JaegerUrl   string
	SentryUrl   string
	KafkaBroker string
	SomeAppId   string
	SomeAppKey  string
}

type envConfig struct {
	DB_URL       string
	JAEGER_URL   string
	SENTRY_URL   string
	KAFKA_BROKER string
}

type flagConfig struct {
	Port       uint64
	SomeAppId  string
	SomeAppKey string
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
	flagConfigData, err := readFlags()
	if err != nil {
		return Config{}, err
	}
	envConfigData, err := readEnv()
	if err != nil {
		return Config{}, err
	}

	return Config{
		Port:        flagConfigData.Port,
		DbUrl:       envConfigData.DB_URL,
		JaegerUrl:   envConfigData.JAEGER_URL,
		SentryUrl:   envConfigData.SENTRY_URL,
		KafkaBroker: envConfigData.KAFKA_BROKER,
		SomeAppId:   flagConfigData.SomeAppId,
		SomeAppKey:  flagConfigData.SomeAppKey,
	}, nil
}

// Чтение флагов
func readFlags() (flagConfig, error) {
	fPort := flag.Uint64("port", 8080, "Port number")
	fSomeAppId := flag.String("some_app_id", "", "Some App ID")
	fSomeAppKey := flag.String("some_app_key", "", "Some App Key")

	flag.Parse()

	if *fPort == 0 {
		return flagConfig{}, errors.New("port can't be empty")
	}
	if *fSomeAppId == "" {
		return flagConfig{}, errors.New("some_app_id can't be empty")
	}
	if *fSomeAppKey == "" {
		return flagConfig{}, errors.New("some_app_key can't be empty")
	}

	return flagConfig{
		Port:       *fPort,
		SomeAppId:  *fSomeAppId,
		SomeAppKey: *fSomeAppKey,
	}, nil
}

// Чтение переменных окружения
func readEnv() (envConfig, error) {
	envConfigData := envConfig{}
	err := envconfig.Process("", &envConfigData)
	if err != nil {
		return envConfig{}, err
	}

	_, err = ConfUrl{url: envConfigData.DB_URL}.decode()
	if err != nil {
		return envConfig{}, errors.New("DB_URL is incorrect")
	}

	_, err = ConfUrl{url: envConfigData.JAEGER_URL}.decode()
	if err != nil {
		return envConfig{}, errors.New("JAEGER_URL is incorrect")
	}

	_, err = ConfUrl{url: envConfigData.SENTRY_URL}.decode()
	if err != nil {
		return envConfig{}, errors.New("SENTRY_URL is incorrect")
	}

	return envConfig{
		DB_URL:       envConfigData.DB_URL,
		JAEGER_URL:   envConfigData.JAEGER_URL,
		SENTRY_URL:   envConfigData.SENTRY_URL,
		KAFKA_BROKER: envConfigData.DB_URL,
	}, nil
}
