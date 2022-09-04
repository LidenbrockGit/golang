package yaml

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port        uint64 `yaml:"port"`
	DbUrl       string `yaml:"db_url"`
	JaegerUrl   string `yaml:"jaeger_url"`
	SentryUrl   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	SomeAppId   string `yaml:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key"`
}

func Read(f string) (Config, error) {
	file, err := os.Open(f)
	if err != nil {
		return Config{}, errors.New("failed to open the file " + f)
	}

	configData := Config{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&configData)
	if err != nil {
		return Config{}, errors.New("failed to decode json file " + f)
	}

	return configData, nil
}
