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
	Data interface{}
}

type Worker struct {
	Tag      string
	Consumer func()
	Logger   *log.Logger
}

var workerChannels map[string]chan WorkerMessage = map[string]chan WorkerMessage{}

var registeredWorkers []Worker = []Worker{}

func SendMessage(tag string, workerMessage WorkerMessage) {
	if c, ok := workerChannels[tag]; ok {
		c <- workerMessage
	} else {
		log.Error("invalid worker message tag '%v'.\n", tag)
	}
}

func RegisterWorker(tag string, consumer func()) {
	registeredWorkers = append(registeredWorkers, Worker{
		Tag:      tag,
		Logger:   log.RootLogger.SubLogger(fmt.Sprintf("[%v worker]", tag)),
		Consumer: consumer,
	})
}

func RegisterAllWorkers() {
	RegisterWorker(WORKER_TELEGRAM, telegramWorkerConsumer)
}

func CreateWorkerChannel(tag string) {
	workerChannels[tag] = make(chan WorkerMessage)
}

func Init() {
	CreateWorkerChannel(WORKER_TELEGRAM)
}

func Start() {
	for _, v := range registeredWorkers {
		go v.Consumer()
	}
}
