package json

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	Port        uint64 `json:"port"`
	DbUrl       string `json:"db_url"`
	JaegerUrl   string `json:"jaeger_url"`
	SentryUrl   string `json:"sentry_url"`
	KafkaBroker string `json:"kafka_broker"`
	SomeAppId   string `json:"some_app_id"`
	SomeAppKey  string `json:"some_app_key"`
}

func Read(f string) (Config, error) {
	file, err := os.Open(f)
	if err != nil {
		return Config{}, errors.New("failed to open the file " + f)
	}

	configData := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configData)
	if err != nil {
		return Config{}, errors.New("failed to decode json file " + f)
	}

	return configData, nil
}
