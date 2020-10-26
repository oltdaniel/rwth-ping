package tasks

import (
	"fmt"
	"time"
)

var messageTask = Task{
	Name:     "Message Task",
	Interval: 5 * time.Second,
	Function: func() {
		fmt.Println("Hello World")
	},
}
