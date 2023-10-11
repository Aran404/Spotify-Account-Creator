package capsolver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func NewSolver(apiKey string, retries int) *Solver {
	return &Solver{
		ApiKey:  apiKey,
		Retries: retries,
		Client: &http.Client{
			Timeout: time.Second * 120,
		},
	}
}

func (s *Solver) CreateTask(task any) error {
	payload := Task{
		ClientKey: s.ApiKey,
		Task:      task,
	}

	encoded, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	resp, err := s.Client.Post("https://api.capsolver.com/createTask", "application/json", bytes.NewBuffer(encoded))

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var Response CreateTaskResponse

	if err := json.NewDecoder(resp.Body).Decode(&Response); err != nil {
		return err
	}

	if Response.ErrorID != 0 {
		return fmt.Errorf("could not create new task, Error: %v", Response.ErrorDescription)
	}

	s.TaskId = Response.TaskID

	return nil
}

func (s *Solver) GetTaskResult() (*CaptchaResponse, error) {
	if !s.Start.IsZero() {
		s.Start = time.Now()
	}

	payload := TaskResult{
		ClientKey: s.ApiKey,
		TaskId:    s.TaskId,
	}

	encoded, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Post("https://api.capsolver.com/getTaskResult", "application/json", bytes.NewBuffer(encoded))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var Response GetTaskResultResponse

	if err := json.NewDecoder(resp.Body).Decode(&Response); err != nil {
		return nil, err
	}

	if Response.ErrorID != 0 {
		return nil, fmt.Errorf("could not get task result, Error: %v", Response.ErrorDescription)
	}

	if Response.Status == "ready" {
		result := new(CaptchaResponse)
		result.CaptchaKey = Response.Solution.GRecaptchaResponse
		result.UserAgent = Response.Solution.UserAgent
		result.Expiry = time.Unix(Response.Solution.ExpireTime, 0)
		result.SolveTime = time.Since(s.Start)

		return result, nil
	}

	s.Retries--

	if s.Retries == 0 {
		return nil, fmt.Errorf("could not get captcha resposne after X retries")
	}

	time.Sleep(time.Second)

	return s.GetTaskResult()
}
