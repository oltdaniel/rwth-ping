package tasks

import (
	"net/http"
	"time"

	"github.com/oltdaniel/rwth-ping/utils"
)

var moodleLandingpageTask = Task{
	Name:     "Moodle Landingpage",
	Interval: 1 * time.Minute,
	Function: func(c *TaskContext) {
		start := time.Now()
		resp, err := http.Get("https://moodle.rwth-aachen.de")
		end := time.Now()
		if err != nil {
			c.Err(err.Error())
			return
		}
		c.Debug("Took %vms.", end.Sub(start).Milliseconds())
		utils.InsertMeasurement(c.TaskSlug(), map[string]interface{}{
			"status_code": resp.StatusCode,
			"url":         "https://moodle.rwth-aachen.de",
		}, map[string]interface{}{
			"response_time": end.Sub(start).Milliseconds(),
		})
	},
}
