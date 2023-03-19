package main

import (
	"os"
	"strings"
)

type configType struct {
	Env            string
	Debug          bool
	AllowedOrigins []string

	TelegramEnable bool
	TelegramId     string
	TelegramSecret string
}

func ConvertEnvToBool(e string) bool {
	return e == "true" || e == "TRUE" || e == "1"
}

func ConvertEnvToStringArr(e string) []string {
	return strings.Split(e, ";")
}

func GetConfig() configType {
	env := os.Getenv("ENV")
	if strings.Contains(os.Args[0], "/_test/") {
		env = "testing"
	}

	//set a default for origins
	origins := ConvertEnvToStringArr(os.Getenv("ALLOWED_ORIGINS"))
	if len(origins) == 0 {
		origins = append(origins, "*")
	}

	return configType{
		Env:            env,
		Debug:          ConvertEnvToBool(os.Getenv("DEBUG")),
		AllowedOrigins: origins,

		TelegramEnable: ConvertEnvToBool(os.Getenv("TELEGRAM_ENABLE")),
		TelegramId:     os.Getenv("TELEGRAM_ID"),
		TelegramSecret: os.Getenv("TELEGRAM_SECRET"),
	}
}
