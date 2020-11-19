package workers

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
		m := <-workerChannels[WORKER_TELEGRAM]
		switch m.Type {
		case WORKER_MESSAGETYPE_TELEGRAM:
			_ = (*m.Data).(TelegramWorkerMessage)
			// create config parser to read target chat ids
			// then parse to telegram message and send
		}
	}
}
