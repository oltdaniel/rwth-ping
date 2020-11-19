package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rsms/go-log"
)

type TelegramSetWebhookRequest struct {
	Url string `json:"url"`
}

func SetTelegramWebhook(host string) {
	log.Debug("Setting webhook to '%v'.", host)
	// construct target url
	url := fmt.Sprintf("https://api.telegram.org/bot%v/setWebhook", CONFIG.Telegram.Token)
	// create request body
	rb := TelegramSetWebhookRequest{
		Url: fmt.Sprintf("%v/webhooks/telegram/%v", host, CONFIG.Telegram.Token),
	}
	// generate json of the message
	v, _ := json.Marshal(rb)
	body := bytes.NewBuffer(v)
	// make the post request
	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		log.Error("Telegram setWebhook failed.")
	}
	// check if error response code
	if true {
		// parse the response
		respText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(string(respText))
			log.Error("Response read failed.")
		}
	}
}
