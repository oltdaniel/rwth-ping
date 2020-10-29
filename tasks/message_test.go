package tasks

import (
	"strings"
	"testing"
	"time"

	"github.com/rsms/go-log"
)

func TestMessageTaskFunction(t *testing.T) {
	// get the buffer of the log output
	buf := testConfigureRootLogger()
	// create a artificial task context
	c := TaskContext{
		Logger:    log.RootLogger,
		StartTime: time.Now(),
	}
	// call the task
	messageTask.Function(&c)
	// wait for all log lines to be printed
	log.Sync()
	// check for text
	if !strings.Contains(buf.String(), "[debug] Hello World.") {
		t.Error("Some error occured. Expected '[debug] Hello World.'")
	}
}
