package main

import (
	"fmt"
	"homework_8/config"
)

func main() {
	configData, err := config.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println("Port: ", configData.Port)
	fmt.Println("DbUrl: ", configData.DbUrl)
	fmt.Println("JaegerUrl: ", configData.JaegerUrl)
	fmt.Println("SentryUrl: ", configData.SentryUrl)
	fmt.Println("KafkaBroker: ", configData.KafkaBroker)
	fmt.Println("SomeAppId: ", configData.SomeAppId)
	fmt.Println("SomeAppKey: ", configData.SomeAppKey)
}
