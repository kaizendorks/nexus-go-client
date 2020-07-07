package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

type ScriptAPI api

// List all stored scripts
//	api endpoint: GET /v1/script
//	responses:
// 		200: successful operation returns a Script slice and nil error
func (a ScriptAPI) List() ([]Script, error) {
	ss := []Script{}
	path := fmt.Sprintf("v1/script")

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return ss, err
	}

	err = json.Unmarshal(resp, &ss)
	return ss, err
}

// Create adds a new script
//	api endpoint: POST /v1/script
// summary: Add a new script
//	parameters:
// 		s: Script object with name, content and type
//	responses:
// 		204: Script was added
// 		410: Script creation is disabled
func (a ScriptAPI) Create(s Script) error {
	path := fmt.Sprintf("v1/script")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(s)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// Get fetches script by name
//	api endpoint: GET /v1/script/{name}
//	parameters:
// 		scriptName: The name of the script.
//	responses:
// 		200: successful operation returns Script object and nil error.
// 		404: No script with the specified name
func (a ScriptAPI) Get(scriptName string) (Script, error) {
	s := Script{}
	path := fmt.Sprintf("v1/script/%s", scriptName)

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(resp, &s)
	return s, err
}

// Update updates the contents of an existing script by name.
//	api endpoint: PUT /v1/script/{name}
//	parameters:
// 		s: Script object
//	responses:
// 		204: Script was updated
// 		404: No script with the specified name
// 		410: Script updating is disabled
func (a ScriptAPI) Update(s Script) error {
	path := fmt.Sprintf("v1/script/%s", s.Name)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(s)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// Delete existing script by name.
//	api endpoint: DELETE /v1/script/{name}
//	parameters:
// 		scriptName: The name of the script to delete.
//	responses:
// 		204: Script was deleted
// 		404: No script with the specified name
func (a ScriptAPI) Delete(scriptName string) error {
	path := fmt.Sprintf("v1/script/%s", scriptName)

	_, err := a.client.sendRequest(http.MethodDelete, path, nil, nil)
	return err
}

// Run executes an existing script
//	api endpoint: PUT /v1/script/{name}/run
//	parameters:
// 		scriptName: the name of the script to execute.
// 		params: an optional key/value map containing string params (use nil for script without params)
//	responses:
// 		200: successful operation returns ScriptResult object and nil error
// 		404: No script with the specified name
// 		500: Script execution failed with exception
func (a ScriptAPI) Run(scriptName string, params map[string]string) (ScriptResult, error) {
	r := ScriptResult{}
	path := fmt.Sprintf("v1/script/%s/run", scriptName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(params)

	resp, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(resp, &r)
	return r, err
}
