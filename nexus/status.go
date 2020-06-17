package nexus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StatusAPI api

type StackTraceElement struct {
	ClassName    string `json:"className"`
	FileName     string `json:"fileName"`
	LineNumber   int32  `json:"lineNumber"`
	MethodName   string `json:"methodName"`
	NativeMethod bool   `json:"nativeMethod"`
}

type Throwable struct {
	Cause            *Throwable           `json:"cause"`
	LocalizedMessage string               `json:"localizedMessage"`
	Message          string               `json:"message"`
	StackTrace       []*StackTraceElement `json:"stackTrace"`
	Suppressed       []*Throwable         `json:"suppressed"`
}

type SystemStatus struct {
	Details   map[string]interface{} `json:"details"`
	Duration  int64                  `json:"duration"`
	Error     *Throwable             `json:"error"`
	Healthy   bool                   `json:"healthy"`
	Message   string                 `json:"message"`
	Time      int64                  `json:"time"`
	Timestamp string                 `json:"timestamp"`
}

type StatusCheckResponse map[string]SystemStatus

// Status health check endpoint that validates if the server can respond to read requests
// api endpoint: GET /v1​/status
// responses:
// 		200: Available to service requests returns nil error
// 		503: Unavailable to service requests
func (a StatusAPI) Status() error {
	path := fmt.Sprintf("v1/status")
	_, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	return err
}

// StatusCheck health check endpoint that returns the results of the system status checks
// api endpoint: GET /v1/status/check
// responses:
// 		200: Successful operation returns StatusCheckResponse and nil error
func (a StatusAPI) StatusCheck() (StatusCheckResponse, error) {
	status := StatusCheckResponse{}
	path := fmt.Sprintf("v1/status/check")

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return status, err
	}

	err = json.Unmarshal(resp, &status)
	return status, err
}

// StatusWritable check endpoint that validates server can respond to read and write requests
// api endpoint: GET /v1​/status/writable
// responses:
// 		200: Available to service requests returns nil error
// 		503: Unavailable to service requests
func (a StatusAPI) StatusWritable() error {
	path := fmt.Sprintf("v1/status/writable")
	_, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	return err
}
