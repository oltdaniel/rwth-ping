package tasks

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/oltdaniel/rwth-ping/utils"
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
	return c.Task.Slug
}

// Register the general task informations.
type Task struct {
	Name     string
	Slug     string
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
	// check if valid
	if !taskSlugRegex.MatchString(t.Slug) {
		panic(fmt.Sprintf("Task slug for '%v' with '%v' not valid.", t.Name, t.Slug))
	}
	// check if enabled
	if !utils.CONFIG.IsEnabled(t.Slug) {
		log.Debug("Task '%v' disabled.", t.Slug)
		return
	}
	// overwrite default interval with custom from config
	t.Interval = utils.CONFIG.GetInterval(t.Slug, t.Interval)
	// init private variables of task
	t.Init()
	// append to registered tasks
	registeredTasks = append(registeredTasks, t)
	// register in scheduler
	clockwerkInstance.Every(t.Interval).Do(t)
	// debug message
	log.Debug("Task '%v' enabled. Every %vseconds.", t.Slug, t.Interval.Seconds())
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
