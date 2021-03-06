package oneview

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/HewlettPackard/oneview-golang/utils"
)

// AssociatedResource associated resource
type AssociatedResource struct {
	ResourceName     utils.Nstring `json:"resourceName,omitempty"`     // "resourceName": "se05, bay 16",
	AssociationType  string        `json:"associationType,omitempty"`  // "associationType": "MANAGED_BY",
	ResourceCateogry string        `json:"resourceCategory,omitempty"` // "resourceCategory": "server-hardware",
	ResourceURI      utils.Nstring `json:"resourceUri,omitempty"`      // "resourceUri": "/rest/server-hardware/30373237-3132-4D32-3235-303930524D57"
}

// TaskState task state
type TaskState int

const (
	T_COMPLETED TaskState = 1 + iota
	T_ERROR
	T_INERRUPTED
	T_KILLED
	T_NEW
	T_PENDING
	T_RUNNING
	T_STARTING
	T_STOPPING
	T_SUSPENDED
	T_TERMINATED
	T_UNKNOWN
	T_WARNING
)

var taskstate = [...]string{
	"Completed",   // Completed Task has been completed.
	"Error",       // Error Task has terminated with an error.
	"Interrupted", // Interrupted Task has been interrupted.
	"Killed",      // Killed Task has been killed.
	"New",         // New Task is new.
	"Pending",     // Pending Task is in pending state.
	"Running",     // Running Task is running.
	"Starting",    // Starting Task is starting.
	"Stopping",    // Stopping Task is stopping.
	"Suspended",   // Suspended Task is suspended.
	"Terminated",  // Terminated Task has been terminated.
	"Unknown",     // Unknown State of task is unknown.
	"Warning",     // Warning Task has terminated with a warning.
}

// String for type
func (ts TaskState) String() string { return taskstate[ts-1] }

// Equal for type
func (ts TaskState) Equal(s string) bool { return (strings.ToUpper(s) == strings.ToUpper(ts.String())) }

// TaskType - task type
type TaskType int

const (
	T_APPLIANCE TaskType = 1 + iota
	T_BACKGROUND
	T_USER
)

var tasktype = [...]string{
	"Applicance", // Appliance Task is appliance initiated and shows in notification panel.
	"Background", // Background Task is appliance initiated and does not show in notification panel.
	"User",       // User Task is user initiated and shows in notification panel.
}

// String return
func (tt TaskType) String() string { return tasktype[tt-1] }

// Equal type
func (tt TaskType) Equal(s string) bool { return (strings.ToUpper(s) == strings.ToUpper(tt.String())) }

// TaskError struct
type TaskError struct {
	Data               map[string]interface{} `json:"data,omitempty"`               // "data":{},
	ErrorCode          string                 `json:"errorCode,omitempty"`          // "errorCode":"MacTypeDiffGlobalMacType",
	Details            string                 `json:"details,omitempty"`            // "details":"",
	NestedErrors       []TaskError            `json:"nestedErrors,omitempty"`       // "nestedErrors":[],
	Message            string                 `json:"message,omitempty"`            // "message":"When macType is not user defined, mac type should be same as the global Mac assignment Virtual."
	ErrorSource        utils.Nstring          `json:"errorSource,omitempty"`        // "errorSource":null,
	RecommendedActions []string               `json:"recommendedActions,omitempty"` // "recommendedActions":["Verify parameters and try again."],
}

// ProgressUpdate - Task Progress Updates
type ProgressUpdate struct {
	TimeStamp    string `json:"timestamp,omitempty"`    // "timestamp":"2015-09-10T22:50:14.250Z",
	StatusUpdate string `json:"statusUpdate,omitempty"` // "statusUpdate":"Apply server settings.",
	ID           int    `json:"id,omitempty"`           // "id":12566
}

// Task structure
type Task struct {
	Type                    string             `json:"type,omitempty"`                    // "type": "TaskResourceV2",
	Data                    TaskData           `json:"data,omitempty"`                    // "data": null,
	Category                string             `json:"category,omitempty"`                // "category": "tasks",
	Hidden                  bool               `json:"hidden,omitempty"`                  // "hidden": false,
	StateReason             string             `json:"stateReason,omitempty"`             // "stateReason": null,
	User                    string             `json:"User,omitempty"`                    // "taskType": "User",
	AssociatedRes           AssociatedResource `json:"associatedResource,omitempty"`      // "associatedResource": { },
	PercentComplete         int                `json:"percentComplete,omitempty"`         // "percentComplete": 0,
	AssociatedTaskURI       utils.Nstring      `json:"associatedTaskUri,omitempty"`       // "associatedTaskUri": null,
	CompletedSteps          int                `json:"completedSteps,omitempty"`          // "completedSteps": 0,
	ComputedPercentComplete int                `json:"computedPercentComplete,omitempty"` //     "computedPercentComplete": 0,
	ExpectedDuration        int                `json:"expectedDuration,omitempty"`        // "expectedDuration": 300,
	ParentTaskURI           utils.Nstring      `json:"parentTaskUri,omitempty"`           // "parentTaskUri": null,
	ProgressUpdates         []ProgressUpdate   `json:"progressUpdates,omitempty"`         // "progressUpdates": [],
	TaskErrors              []TaskError        `json:"taskErrors,omitempty"`              // "taskErrors": [],
	TaskOutput              []string           `json:"taskOutput,omitempty"`              // "taskOutput": [],
	TaskState               string             `json:"taskState,omitempty"`               // "taskState": "New",
	TaskStatus              string             `json:"taskStatus,omitempty"`              // "taskStatus": "Power off Server: se05, bay 16",
	TaskType                string             `json:"taskType,omitempty"`
	TotalSteps              int                `json:"totalSteps,omitempty"`    // "totalSteps": 0,
	UserInitiated           bool               `json:"userInitiated,omitempty"` // "userInitiated": true,
	Name                    string             `json:"name,omitempty"`          // "name": "Power off",
	Owner                   string             `json:"owner,omitempty"`         // "owner": "wenlock",
	ETAG                    string             `json:"eTag,omitempty"`          // "eTag": "0",
	Created                 string             `json:"created,omitempty"`       // "created": "2015-09-07T03:25:54.844Z",
	Modified                string             `json:"modified,omitempty"`      // "modified": "2015-09-07T03:25:54.844Z",
	URI                     string             `json:"uri,omitempty"`           // "uri": "/rest/tasks/145F808A-A8DD-4E1B-8C86-C2379C97B3B2"
	TaskIsDone              bool               // when true, task are done
	Timeout                 int                // time before timeout on Executor
	WaitTime                time.Duration      // time between task checks
	Client                  *CLIOVClient
}

// TaskServer Example:
// {"name":"se05, bay 14", "uri":"/rest/server-hardware/30373237-3132-4D32-3235-303930524D52"}
type TaskServer struct {
	Name string `json:"name,omitempty"` // "Name to server
	URI  string `json:"uri,omitempty"`  // "URI to server
}

type TaskData struct {
	TaskCategory string `json:"task-category,omitempty"`
}

// NewProfileTask - Create New Task
func NewTask(c *CLIOVClient) *Task {
	return &Task{
		TaskIsDone: false,
		Client:     c,
		URI:        "",
		Name:       "",
		Owner:      "",
		Timeout:    270,  // default 45min
		WaitTime:   2000, // default 2sec, impacts Timeout
	}
}

// ResetTask - reset the power task back to off
func (t *Task) ResetTask() {
	t.TaskIsDone = false
	t.URI = ""
	t.Name = ""
	t.Owner = ""
}

// GetCurrentTaskStatus - Get the current status
func (t *Task) GetCurrentTaskStatus() error {
	log.Print("[DEBUG] Working on getting current task status")
	var (
		uri = t.URI
	)
	if uri != "" {
		log.Print("[DEBUG]", uri)
		data, err := t.Client.SendHTTPRequest("GET", uri, nil)
		if err != nil {
			return err
		}
		log.Printf("[DEBUG] data: %s", data)
		if err := json.Unmarshal([]byte(data), &t); err != nil {
			return err
		}
	} else {
		log.Printf("Unable to get current task, no URI found")
	}
	if len(t.TaskErrors) > 0 {
		var errmsg string
		errmsg = ""
		for _, te := range t.TaskErrors {
			errmsg += te.Message + " \n" + strings.Join(te.RecommendedActions, " ")
		}
		return errors.New(errmsg)
	}
	return nil
}

// GetLastStatusUpdate - get last detail updates from task
func (t *Task) GetLastStatusUpdate() string {
	if len(t.ProgressUpdates) > 0 {
		lastupdate := len(t.ProgressUpdates) - 1
		// sanatize a little by removing json
		message := utils.StringRemoveJSON(t.ProgressUpdates[lastupdate].StatusUpdate)
		// parse out server name
		servernamejson := utils.StringGetJSON(t.ProgressUpdates[lastupdate].StatusUpdate)
		var ts *TaskServer
		if err := json.Unmarshal([]byte(servernamejson), &ts); err == nil {
			message += ts.Name
		}
		return t.TaskStatus + ", " + message
	}
	return t.TaskStatus
}

// Wait - wait on task to complete
func (t *Task) Wait() error {

	for t.PercentComplete != 100 {
		t.checkStatus()
		time.Sleep(time.Millisecond * t.WaitTime)
	}

	fmt.Printf("*** Task Final -State: %v, -Status: %v\n", t.TaskState, t.TaskStatus)
	if len(t.TaskErrors) != 0 {
		fmt.Printf("Error Code: %v\n", t.TaskErrors[0].ErrorCode)
		fmt.Printf("Message: %v\n", t.TaskErrors[0].Message)
		fmt.Printf("Recommendation: %v\n", t.TaskErrors[0].RecommendedActions)
		fmt.Printf("Details: %v\n", t.TaskErrors[0].Details)

		os.Exit(1)
	}

	return nil
}

func (t *Task) checkStatus() error {
	data, err := t.Client.SendHTTPRequest("GET", t.URI, nil)
	if err != nil {
		fmt.Printf("OVCLI task wait failure to get task data from task URI: %v", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(data, &t); err != nil {
		fmt.Printf("OVCLI task wait failure to get task json data decoded: %v", err)
		os.Exit(1)
	}

	fmt.Printf("state: %v, status: %v\n", t.TaskState, t.TaskStatus)
	fmt.Printf("expectedDuration: %v, completedSteps: %v, percentComplete: %v, computedPercentCompute: %v\n\n", t.ExpectedDuration, t.CompletedSteps, t.PercentComplete, t.ComputedPercentComplete)

	return nil
}

// var (
// 	currenttime int
// )
// log.Printf("[DEBUG] task : %+v", t)
// if t.Timeout < t.ExpectedDuration {
// 	t.Timeout = t.ExpectedDuration
// 	log.Printf("[DEBUG] OVCLI increase task timeout to the replied ExpectedDuration %d", t.Timeout)
// }
// log.Printf("[DEBUG] task timeout is : %d", t.Timeout)
// for !t.TaskIsDone && (currenttime < t.Timeout) {
// 	if err := t.GetCurrentTaskStatus(); err != nil {
// 		t.TaskIsDone = true
// 		return err
// 	}
// 	if t.URI != "" && T_COMPLETED.Equal(t.TaskState) {
// 		t.TaskIsDone = true
// 	}
// 	if t.URI != "" {
// 		log.Printf("[DEBUG] Waiting for task to complete, for %s ", t.Name)
// 		log.Printf("[DEBUG] Waiting on, %s, %d%%, %s, %d, %d", t.Name, t.ComputedPercentComplete, t.GetLastStatusUpdate(), currenttime, t.ExpectedDuration)
// 	} else {
// 		log.Printf("[DEBUG] Waiting on task creation.")
// 	}

// 	// wait time before next check
// 	time.Sleep(time.Millisecond * (1000 * t.WaitTime)) // wait 10sec before checking the status again
// 	currenttime++
// 	if t.Timeout < t.ExpectedDuration {
// 		t.Timeout = t.ExpectedDuration
// 	}
// }
// if currenttime > t.Timeout {
// 	log.Printf("[DEBUG] Task timed out, %d.", currenttime)
// }

// if t.Name != "" {
// 	log.Printf("[DEBUG] Task, %s, completed", t.Name)
// }
// return nil
// }
