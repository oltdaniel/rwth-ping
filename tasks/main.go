package tasks

import (
	"fmt"
	"time"

	"github.com/onatm/clockwerk"
)

// Register the general task informations.
type Task struct {
	Name     string
	Interval time.Duration
	Function func()
}

// Wrap the task function in a warpper.
func (t Task) Run() {
	fmt.Printf("Running: '%v'.\n", t.Name)
	t.Function()
}

// Keep track of the registered tracks.
var registeredTasks = []Task{}

// The scheduler instance.
var clockwerkInstance = clockwerk.New()

// Register all known tasks in the system.
func RegisterAllTasks() {
	RegisterTask(messageTask)
}

// Register a single task in the scheduler and the registered Task array.
func RegisterTask(t Task) {
	registeredTasks = append(registeredTasks, t)
	clockwerkInstance.Every(t.Interval).Do(t)
}

// Start the scheduler.
func Start() {
	clockwerkInstance.Start()
}
