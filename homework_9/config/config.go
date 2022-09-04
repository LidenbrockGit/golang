package config

import (
	"errors"
	"flag"
	"github.com/kelseyhightower/envconfig"
	configJSON "homework_9/config/json"
	configYAML "homework_9/config/yaml"
	"net/url"
	"path/filepath"
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

var fPort = flag.Uint64("port", 8080, "Port number")
var fSomeAppId = flag.String("some_app_id", "", "Some App ID")
var fSomeAppKey = flag.String("some_app_key", "", "Some App Key")

func (u ConfUrl) decode() (string, error) {
	var response string
	parseResult, err := url.Parse(u.url)
	if parseResult != nil {
		response = parseResult.String()
	}
	return response, err
}

// Read возвращает данные конфигурации;
// чтение данных происходит через флаги, переменные окружения или
// переданный в аргументе файл
func Read() (Config, error) {
	flag.Parse()
	configData := Config{}

	err := readFile(&configData)
	if err != nil {
		return Config{}, err
	}

	err = readFlags(&configData)
	if err != nil {
		return Config{}, err
	}

	err = readEnv(&configData)
	if err != nil {
		return Config{}, err
	}

	return configData, nil
}

// Чтение файлов
func readFile(confData *Config) error {
	if flag.NArg() == 0 {
		return nil
	}

	fileName := flag.Arg(0)
	fileExt := filepath.Ext(fileName)

	switch fileExt {
	case ".json":
		rawConfigData, err := configJSON.Read(fileName)
		if err != nil {
			return err
		}
		confData.Port = rawConfigData.Port
		confData.DbUrl = rawConfigData.DbUrl
		confData.JaegerUrl = rawConfigData.JaegerUrl
		confData.SentryUrl = rawConfigData.SentryUrl
		confData.KafkaBroker = rawConfigData.KafkaBroker
		confData.SomeAppId = rawConfigData.SomeAppId
		confData.SomeAppKey = rawConfigData.SomeAppKey

	case ".yaml":
		rawConfigData, err := configYAML.Read(fileName)
		if err != nil {
			return err
		}
		confData.Port = rawConfigData.Port
		confData.DbUrl = rawConfigData.DbUrl
		confData.JaegerUrl = rawConfigData.JaegerUrl
		confData.SentryUrl = rawConfigData.SentryUrl
		confData.KafkaBroker = rawConfigData.KafkaBroker
		confData.SomeAppId = rawConfigData.SomeAppId
		confData.SomeAppKey = rawConfigData.SomeAppKey

	default:
		return errors.New("Unknown file extension " + fileExt)
	}

	err := validRawConfigData(confData)
	if err != nil {
		return err
	}

	return nil
}

// Чтение флагов
func readFlags(confData *Config) error {
	if flag.NArg() != 0 {
		return nil
	}

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
	if flag.NArg() != 0 {
		return nil
	}

	err := envconfig.Process("", confData)
	if err != nil {
		return err
	}

	err = validRawConfigData(confData)
	if err != nil {
		return err
	}

	return nil
}

// Валидация данных конфигурации
func validRawConfigData(c *Config) error {
	_, err := ConfUrl{url: c.DbUrl}.decode()
	if err != nil {
		return errors.New("DB_URL is incorrect")
	}

	_, err = ConfUrl{url: c.JaegerUrl}.decode()
	if err != nil {
		return errors.New("JAEGER_URL is incorrect")
	}

	_, err = ConfUrl{url: c.SentryUrl}.decode()
	if err != nil {
		return errors.New("SENTRY_URL is incorrect")
	}

	return nil
}
