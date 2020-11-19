package tasks

import (
	"fmt"
	"net/http"
	"time"

	"github.com/oltdaniel/rwth-ping/utils"
	"github.com/oltdaniel/rwth-ping/workers"
)

var moodleLandingpageTask = Task{
	Name:     "Moodle Landingpage",
	Interval: 1 * time.Second,
	Function: func(c *TaskContext) {
		// record time
		start := time.Now()
		// make request
		resp, err := http.Get("https://moodle.rwth-aachen.de")
		// record end time
		end := time.Now()
		// check for request error
		if err != nil {
			c.Err(err.Error())
			return
		}
		// debug message on response time
		c.Debug("Took %vms.", end.Sub(start).Milliseconds())
		// record mesaurement
		utils.InsertMeasurement(c.TaskSlug(), map[string]interface{}{
			"status_code": resp.StatusCode,
			"url":         "https://moodle.rwth-aachen.de",
		}, map[string]interface{}{
			"response_time": end.Sub(start).Milliseconds(),
		})
		// check if we have an error here
		if resp.StatusCode >= 400 {
			// send message to the relevant workers
			workers.SendMessage(workers.WORKER_TELEGRAM, workers.WorkerMessage{
				Type: workers.WORKER_MESSAGETYPE_TELEGRAM,
				Data: workers.TelegramWorkerMessage{
					TargetGroup: "moodle",
					Text: fmt.Sprintf(
						"☢️ Moodle Error ☢️\n\n*Status*: %v\\.\n*Took*: %vms\\.",
						resp.StatusCode,
						end.Sub(start).Milliseconds(),
					),
					ParseMode: "MarkdownV2",
				},
			})
		}
	},
}
