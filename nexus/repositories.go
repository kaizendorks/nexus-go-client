package nexus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RepositoriesAPI api

type RepositoryV1 struct {
	Attributes map[string]interface{} `json:"attributes"`
	Format     string                 `json:"format"`
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	URL        string                 `json:"url"`
}

// List return a slice of RepositoryV1 objects
// api endpoint: GET ​/v1​/repositories
// responses:
// 		200: successful operation returns slice of RepositoryV1 and nil error
func (a RepositoriesAPI) List() ([]RepositoryV1, error) {
	rr := []RepositoryV1{}
	path := fmt.Sprintf("v1/repositories")

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return rr, err
	}
	err = json.Unmarshal(resp, &rr)
	return rr, err
}
