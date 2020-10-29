package tasks

import (
	"net/http"
	"time"
)

var moodleLandingpageTask = Task{
	Name:     "Moodle landingpage",
	Interval: 15 * time.Second,
	Function: func(c *TaskContext) {
		start := time.Now()
		_, err := http.Get("https://moodle.rwth-aachen.de")
		end := time.Now()
		if err != nil {
			c.Err(err.Error())
		}
		c.Debug("Took %vms.", end.Sub(start).Milliseconds())
	},
}
