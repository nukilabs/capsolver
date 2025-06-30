package capsolver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	ApiURL     = "https://api.capsolver.com"
	AppID      = "D3119ABC-FF91-42EF-9F18-C4CE4B259E52"
	MaxRetries = 120
)

type Session struct {
	key    string
	client *http.Client
}

func New(key string) *Session {
	return &Session{
		key:    key,
		client: &http.Client{},
	}
}

func (s *Session) post(endpoint string, payload Payload) (*Result, error) {
	payload.ClientKey = s.key
	payload.AppID = AppID

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, ApiURL+endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result Result
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *Session) createTask(task any) (*Result, error) {
	return s.post("/createTask", Payload{Task: task})
}

func (s *Session) getTaskResult(taskID string) (*Result, error) {
	return s.post("/getTaskResult", Payload{TaskID: taskID})
}

func (s *Session) Solve(task any) (*Result, error) {
	res, err := s.createTask(task)
	if err != nil {
		return nil, err
	}
	if res.Error.ID == 1 {
		return nil, res.Error
	}
	for i := 0; res.Status != StatusReady && i < MaxRetries; i++ {
		time.Sleep(3 * time.Second)

		res, err = s.getTaskResult(res.TaskId)
		if err != nil {
			return nil, err
		}
		if res.Error.ID == 1 {
			return nil, res.Error
		}
	}
	return res, nil
}
