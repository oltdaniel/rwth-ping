package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oltdaniel/rwth-ping/tasks"
)

var DEBUG = (os.Getenv("D") == "1")

func main() {
	// construct new web instance
	s := gin.New()

	// Assign middlewares
	if DEBUG {
		// use logger
		s.Use(gin.Logger())
	} else {
		// recover from hard failures
		s.Use(gin.Recovery())
	}

	// register start
	s.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, world")
	})

	// initialize all components
	tasks.Init()
	// register all the tasks we know
	tasks.RegisterAllTasks()
	// start the scheduler
	tasks.Start()

	// run the gin server
	log.Fatal(s.Run(":4001"))
}
