package nexus

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

type StatusAPI api

// Status health check endpoint that validates if the server can respond to read requests
//	api endpoint: GET /v1​/status
//	responses:
// 		200: Available to service requests returns nil error
// 		503: Unavailable to service requests
func (a StatusAPI) Status() error {
	path := fmt.Sprintf("v1/status")
	_, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	return err
}

// StatusCheck health check endpoint that returns the results of the system status checks
//	api endpoint: GET /v1/status/check
//	responses:
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
//	api endpoint: GET /v1​/status/writable
//	responses:
// 		200: Available to service requests returns nil error
// 		503: Unavailable to service requests
func (a StatusAPI) StatusWritable() error {
	path := fmt.Sprintf("v1/status/writable")
	_, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	return err
}
