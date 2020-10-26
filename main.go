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
	s := gin.New()

	// Assign middlewares
	if DEBUG {
		s.Use(gin.Logger())
	} else {
		s.Use(gin.Recovery())
	}

	// register start
	s.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, world")
	})

	// register all the tasks we know
	tasks.RegisterAllTasks()
	// start the scheduler
	tasks.Start()

	// run the gin server
	log.Fatal(s.Run(":4001"))
}
