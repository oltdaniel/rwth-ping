package tasks

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/onatm/clockwerk"
	"github.com/rsms/go-log"
)

// check if the debug flag is set
var DEBUG = (os.Getenv("D") == "1")

// regex to create task slug
var taskSlugRegex = regexp.MustCompile("[^a-z0-9]")

// The context a execution is runnung in
type TaskContext struct {
	Task      *Task
	Logger    *log.Logger
	StartTime time.Time
}

// Alias the debug function
func (c *TaskContext) Debug(message string, v ...interface{}) {
	c.Logger.Debugf(message, v...)
}

// Alias the info function
func (c *TaskContext) Info(message string, v ...interface{}) {
	c.Logger.Debugf(message, v...)
}

// Alias the warning function
func (c *TaskContext) Warn(message string, v ...interface{}) {
	c.Logger.Warningf(message, v...)
}

// Alias the error function
func (c *TaskContext) Err(message string, v ...interface{}) {
	c.Logger.Errorf(message, v...)
}

// Create a lower case, non special symbol, no space string
func (c *TaskContext) TaskSlug() string {
	return taskSlugRegex.ReplaceAllString(strings.ReplaceAll(strings.ToLower(c.Task.Name), " ", "-"), "")
}

// Register the general task informations.
type Task struct {
	Name     string
	Interval time.Duration
	Function func(*TaskContext)
	logger   *log.Logger
}

// Initialize all task values.
func (t *Task) Init() {
	t.logger = log.RootLogger.SubLogger(fmt.Sprintf("[%v]", t.Name))
}

// Wrap the task function in a warpper.
func (t Task) Run() {
	t.logger.Debug("Starting.")
	// construct context for this execution
	context := TaskContext{
		Task:      &t,
		Logger:    t.logger.SubLogger("[exec]"),
		StartTime: time.Now(),
	}
	// execute function (will be non-blocking due to the scheduler)
	t.Function(&context)
}

// Keep track of the registered tracks.
var registeredTasks = []Task{}

// The scheduler instance.
var clockwerkInstance = clockwerk.New()

// Register all known tasks in the system.
func RegisterAllTasks() {
	RegisterTask(messageTask)
	RegisterTask(moodleLandingpageTask)
}

// Register a single task in the scheduler and the registered Task array.
func RegisterTask(t Task) {
	// init private variables of task
	t.Init()
	// append to registered tasks
	registeredTasks = append(registeredTasks, t)
	// register in scheduler
	clockwerkInstance.Every(t.Interval).Do(t)
}

// Start the scheduler.
func Start() {
	// start the scheduler
	clockwerkInstance.Start()
}

// Initialize all required components.
func Init() {
	// check for debug flag
	if DEBUG {
		// make debug messages visible
		log.RootLogger.Level = log.LevelDebug
	}
	// enable extended timeformat
	log.RootLogger.EnableFeatures(log.FMicroseconds)
}
