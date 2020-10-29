package tasks

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/rsms/go-log"
)

// src: https://github.com/rsms/go-log/blob/master/log_test.go#L18
func testConfigureRootLogger() *bytes.Buffer {
	// make some changes
	log.RootLogger.DisableFeatures(log.FColor)
	log.RootLogger.EnableFeatures(log.FMicroseconds)
	log.RootLogger.Level = log.LevelDebug
	// some buffer stuff
	w := &bytes.Buffer{}
	if testing.Verbose() {
		log.RootLogger.SetWriter(io.MultiWriter(w, os.Stdout))
	} else {
		log.RootLogger.SetWriter(w)
	}
	return w
}
