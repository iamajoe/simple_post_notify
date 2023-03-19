package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func sendTelegramMsg(id string, secret string, msg string) (bool, error) {
	// do not proceed with send to telegram if we are running tests
	if GetConfig().Env == "testing" {
		return true, nil
	}

	var err error
	var response *http.Response

	rawUrl := fmt.Sprintf("https://api.telegram.org/bot%s", secret)
	url := fmt.Sprintf("%s/sendMessage", rawUrl)
	body, _ := json.Marshal(map[string]string{
		"chat_id": id,
		"text":    msg,
	})
	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	log.Println("Telegram: Message sent")

	return true, nil
}
