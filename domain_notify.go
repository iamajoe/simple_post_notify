package main

import (
	"errors"
	"net/http"
)

// TODO: set tests

func sendNotifyMsg(msg string, mod string) (bool, error) {
	switch mod {
	case "telegram":
		if GetConfig().TelegramEnable {
			return sendTelegramMsg(GetConfig().TelegramId, GetConfig().TelegramSecret, msg)
		}
	}

	return false, NewError(http.StatusBadRequest, errors.New("module not found"))
}
