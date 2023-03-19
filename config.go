package main

import (
	"os"
	"strings"
)

type configType struct {
	Env   string
	Debug bool

	TelegramEnable bool
	TelegramId     string
	TelegramSecret string
}

func ConvertEnvToBool(e string) bool {
	return e == "true" || e == "TRUE" || e == "1"
}

func GetConfig() configType {
	env := os.Getenv("ENV")
	if strings.Contains(os.Args[0], "/_test/") {
		env = "testing"
	}

	return configType{
		Env:            env,
		Debug:          ConvertEnvToBool(os.Getenv("DEBUG")),

		TelegramEnable: ConvertEnvToBool(os.Getenv("TELEGRAM_ENABLE")),
		TelegramId:     os.Getenv("TELEGRAM_ID"),
		TelegramSecret: os.Getenv("TELEGRAM_SECRET"),
	}
}
