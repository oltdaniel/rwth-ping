package workers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/oltdaniel/rwth-ping/utils"
	"github.com/rsms/go-log"
)

type TelegramMessage struct {
	ChatId    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type TelegramWorkerMessage struct {
	TargetGroup string
	Text        string
	ParseMode   string
}

const WORKER_MESSAGETYPE_TELEGRAM = "worker_messsagetype_telegram"

func telegramWorkerConsumer() {
	for {
		// wait for next message
		m := <-workerChannels[WORKER_TELEGRAM]
		// cast interface depending on type
		switch m.Type {
		case WORKER_MESSAGETYPE_TELEGRAM:
			// parse message in correct format
			twm := m.Data.(TelegramWorkerMessage)
			// check if the wanted target group exists
			if g, ok := utils.CONFIG.Telegram.TargetGroups[twm.TargetGroup]; ok {
				// for each target group
				for _, v := range g {
					// construct telegram message
					tm := TelegramMessage{
						ChatId:    v,
						Text:      twm.Text,
						ParseMode: twm.ParseMode,
					}
					sendTelegramMessage(tm)
				}
			}
		}
	}
}

func sendTelegramMessage(tm TelegramMessage) {
	log.Debug("Sending telegram message to '%v'.", tm.ChatId)
	// construct target url
	url := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", utils.CONFIG.Telegram.Token)
	// generate json of the message
	v, _ := json.Marshal(tm)
	body := bytes.NewBuffer(v)
	// make the post request
	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		log.Error("Telegram message failed.")
	}
	// check if error response code
	if resp.StatusCode >= 400 {
		// parse the response
		respText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(string(respText))
			log.Error("Response read failed.")
		}
	}
}
