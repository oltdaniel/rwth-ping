package workers

import (
	"fmt"

	"github.com/rsms/go-log"
)

const (
	WORKER_TELEGRAM = "telegram_worker"
)

type WorkerMessage struct {
	Type string
	Data *interface{}
}

type Worker struct {
	Tag      string
	Consumer func()
	Logger   *log.Logger
}

var workerChannels map[string]chan WorkerMessage = map[string]chan WorkerMessage{
	WORKER_TELEGRAM: make(chan WorkerMessage),
}

var registeredWorkers []Worker = []Worker{}

func BroadcastMessage(tag string, workerMessage WorkerMessage) {
	if c := workerChannels[tag]; c != nil {
		c <- workerMessage
	}
	log.Error("invalid worker message tag '%v'.\n", tag)
}

func RegisterWorker(tag string, consumer func()) {
	registeredWorkers = append(registeredWorkers, Worker{
		Tag: tag, Logger: log.RootLogger.SubLogger(fmt.Sprintf("[%v worker]", tag)),
	})
}

func StartWorkers() {
	for _, v := range registeredWorkers {
		go v.Consumer()
	}
}
