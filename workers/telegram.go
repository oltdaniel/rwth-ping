package workers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

const WORKER_MESSAGETYPE_TELEGRAM_SEND = "worker_messagetype_telegram_send"
const WORKER_TELEGRAM_SEND = "worker_telegram_send"

func telegramWorkerConsumer() {
	for {
		// wait for next message
		m := <-workerChannels[WORKER_TELEGRAM_SEND]
		// cast interface depending on type
		switch m.Type {
		case WORKER_MESSAGETYPE_TELEGRAM_SEND:
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

const WORKER_MESSAGETYPE_TELEGRAM_RECEIVE_PRIVATE = "worker_messagetype_telegram_receive_private"
const WORKER_MESSAGETYPE_TELEGRAM_RECEIVE_GROUP = "worker_messagetype_telegram_receive_group"
const WORKER_TELEGRAM_RECEIVE = "worker_telegram_receive"

type TelegramWorkerMessageReceive = TelegramUpdate

func telegramWorkerRespondToPrivateMessages() {
	for {
		// wait for next message
		m := <-workerChannels[WORKER_TELEGRAM_RECEIVE]
		// cast interface depending on type
		switch m.Type {
		case WORKER_MESSAGETYPE_TELEGRAM_RECEIVE_PRIVATE:
			// parse message in correct format
			twm := m.Data.(TelegramWorkerMessageReceive)
			// respond to private chat
			tm := TelegramMessage{
				ChatId:    strconv.Itoa(twm.Message.Chat.Id),
				Text:      "Hi, I don't support this\\.",
				ParseMode: "MarkdownV2",
			}
			sendTelegramMessage(tm)
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

type TelegramUpdate struct {
	UpdateId int `json:"update_id"`
	Message  struct {
		MessageId int `json:"message_id"`
		Chat      struct {
			Id       int    `json:"id"`
			Username string `json:"username"`
			Type     string `json:"type"`
		} `json:"chat"`
		From struct {
			Id       int    `json:"id"`
			Username string `json:"username"`
		} `json:"from"`
		Text string `json:"text"`
		Date int    `json:"date"`
	}
}

func TelegramWebhookHandler(c *gin.Context) {
	// verify provided token from path parameter
	givenToken := c.Param("token")
	if givenToken != utils.CONFIG.Telegram.Token {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "hello hacker, bye hacker",
		})
		return
	}
	// parse body to struct
	var tu TelegramUpdate
	if err := c.BindJSON(&tu); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err,
		})
		return
	}
	// done response
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "done",
	})
	// send message to workers depending on type
	if tu.Message.Chat.Type == "private" {
		// send private message
		SendMessage(WORKER_TELEGRAM_RECEIVE, WorkerMessage{
			Type: WORKER_MESSAGETYPE_TELEGRAM_RECEIVE_PRIVATE,
			Data: tu,
		})
	}
}
