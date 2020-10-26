package tasks

import (
	"fmt"
	"time"

	"github.com/onatm/clockwerk"
)

type Task struct {
	Name     string
	Interval time.Duration
	Function func()
}

func (t Task) Run() {
	fmt.Printf("Running: '%v'.\n", t.Name)
	t.Function()
}

var registeredTasks = []Task{}
var clockwerkInstance = clockwerk.New()

func RegisterAllTasks() {
	RegisterTask(messageTask)
}

func RegisterTask(t Task) {
	registeredTasks = append(registeredTasks, t)
	clockwerkInstance.Every(t.Interval).Do(t)
}

func Start() {
	clockwerkInstance.Start()
}
