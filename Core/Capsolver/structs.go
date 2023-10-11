package capsolver

import (
	"net/http"
	"time"
)

type Solver struct {
	ApiKey  string
	Client  *http.Client
	Retries int
	TaskId  string
	Start   time.Time
}

type Task struct {
	ClientKey string `json:"clientKey"`
	Task      any    `json:"task"`
}

type TaskResult struct {
	ClientKey string `json:"clientKey"`
	TaskId    string `json:"taskId"`
}

type CreateTaskResponse struct {
	ErrorID          int    `json:"errorId"`
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	TaskID           string `json:"taskId"`
}

type GetTaskResultResponse struct {
	ErrorID          int `json:"errorId"`
	ErrorCode        any `json:"errorCode"`
	ErrorDescription any `json:"errorDescription"`
	Solution         struct {
		UserAgent          string `json:"userAgent"`
		ExpireTime         int64  `json:"expireTime"`
		GRecaptchaResponse string `json:"gRecaptchaResponse"`
	} `json:"solution"`
	Status string `json:"status"`
}

type CaptchaResponse struct {
	UserAgent  string
	CaptchaKey string
	TaskId     string
	Expiry     time.Time
	SolveTime  time.Duration
}
