package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GolangGroupRepository struct {
	Group *Group `json:"group"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type GolangProxyRepository struct {
	Cleanup    *Cleanup    `json:"cleanup,omitempty"`
	HTTPClient *HTTPClient `json:"httpClient"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	NegativeCache *NegativeCache `json:"negativeCache"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Proxy       *Proxy   `json:"proxy"`
	RoutingRule string   `json:"routingRule,omitempty"`
	Storage     *Storage `json:"storage"`
}

// CreateGolangGroup creates Golang group repository
// endpoint: POST ​/beta​/repositories​/go/group
// parameters:
// 		r:
// 			description: GolangGroupRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateGolangGroup(r GolangGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/go/group")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateGolangGroup updates Golang group repository
// endpoint: PUT ​/beta​/repositories​/go/group​/{repositoryName}
// parameters:
// 		r:
// 			description: GolangGroupRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
//    401: Authentication required
//    403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) UpdateGolangGroup(repositoryName string, r GolangGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/go/group/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateGolangProxy creates Golang proxy repository
// endpoint: POST ​/beta​/repositories​/go/proxy
// parameters:
// 		r:
// 			description: GolangProxyRepository config
// responses:
//    201: Repository created
//    401: Authentication required
//    403: Insufficient permissions
func (a RepositoryManagementAPI) CreateGolangProxy(r GolangProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/go/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateGolangProxy updates Golang proxy repository
// endpoint: PUT ​/beta​/repositories​/go/proxy​/{repositoryName}
// parameters:
// 		r:
// 			description: GolangProxyRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
//    401: Authentication required
//    403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) UpdateGolangProxy(repositoryName string, r GolangProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/go/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
