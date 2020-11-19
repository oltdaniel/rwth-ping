package main

import (
	stdlog "log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oltdaniel/rwth-ping/tasks"
	"github.com/oltdaniel/rwth-ping/utils"
	"github.com/oltdaniel/rwth-ping/workers"
	"github.com/rsms/go-log"
)

var DEBUG = (os.Getenv("D") == "1")

func main() {
	// construct new web instance
	s := gin.New()

	// Assign middlewares
	if DEBUG {
		// use logger
		s.Use(gin.Logger())
		log.RootLogger.Level = log.LevelDebug
	} else {
		// recover from hard failures
		s.Use(gin.Recovery())
	}

	// register start
	s.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, world")
	})

	// hook telegram webhook listener
	s.POST("/webhooks/telegram/:token", workers.TelegramWebhookHandler)
	if host := os.Getenv("HOST"); host != "" {
		utils.SetTelegramWebhook(host)
	}

	// initialize workers toolbox
	workers.Init()
	// register all exisiting workers
	workers.RegisterAllWorkers()
	// start all workers
	workers.Start()

	// initialize all tasks components
	tasks.Init()
	// register all the tasks we know
	tasks.RegisterAllTasks()
	// start the scheduler
	tasks.Start()

	// run the gin server
	stdlog.Fatal(s.Run(":4001"))
}
