package capsolver

import "encoding/json"

// Payload represents the request payload for the API.
type Payload struct {
	ClientKey string `json:"clientKey"`
	AppID     string `json:"appId,omitempty"`
	Task      any    `json:"task,omitempty"`
	TaskID    string `json:"taskId,omitempty"`
}

// Status represents the status of a task.
type Status string

const (
	// StatusReady indicates the identification is completed and the solution is ready.
	StatusReady Status = "ready"
	// StatusIdle indicates the task is idle and waiting for processing.
	StatusIdle Status = "idle"
	// StatusProcessing indicates the task is currently being processed.
	StatusProcessing Status = "processing"
)

// Result represents the response from the API.
type Result struct {
	Status   Status          `json:"status"`
	Solution json.RawMessage `json:"solution"`
	TaskId   string          `json:"taskId"`
	Error
}

// Unmarshal unmarshals the solution into the appropriate type based on the task type.
func (r *Result) Unmarshal(v any) error {
	return json.Unmarshal(r.Solution, v)
}
