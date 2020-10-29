package tasks

import (
	"time"
)

var messageTask = Task{
	Name:     "Message Task",
	Interval: 5 * time.Second,
	Function: func(c *TaskContext) {
		c.Debug("Hello World.")
	},
}
