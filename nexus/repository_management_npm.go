package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type NPMGroupRepository struct {
	Group *Group `json:"group"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type NPMHostedRepository struct {
	Cleanup *Cleanup `json:"cleanup,omitempty"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type NPMProxyRepository struct {
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

// CreateNPMGroup creates NPM group repository
// endpoint: POST ​/beta​/repositories​/npm​/group
// parameters:
// 		r:
// 			description: NPMGroupRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateNPMGroup(r NPMGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/npm/group")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateNPMGroup updates NPM group repository
// endpoint: PUT ​/beta​/repositories​/npm​/group​/{repositoryName}
// parameters:
// 		r:
// 			description: NPMGroupRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
//    401: Authentication required
//    403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) UpdateNPMGroup(repositoryName string, r NPMGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/npm/group/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateNPMHosted create NPM hosted repository
// endpoint: POST ​/beta​/repositories​/npm​/hosted
// parameters:
// 		r:
// 			description: NPMHostedRepository config
// responses:
//    201: Repository created
//    401: Authentication required
//    403: Insufficient permissions
func (a RepositoryManagementAPI) CreateNPMHosted(r NPMHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/npm/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateNPMHosted updates NPM hosted repository
// endpoint: PUT ​/beta​/repositories​/npm​/hosted​/{repositoryName}
// parameters:
// 		r:
// 			description: NPMHostedRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
//    401: Authentication required
//    403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) UpdateNPMHosted(repositoryName string, r NPMHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/npm/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateNPMProxy creates NPM proxy repository
// endpoint: POST ​/beta​/repositories​/npm​/proxy
// parameters:
// 		r:
// 			description: NPMProxyRepository config
// responses:
//    201: Repository created
//    401: Authentication required
//    403: Insufficient permissions
func (a RepositoryManagementAPI) CreateNPMProxy(r NPMProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/npm/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateNPMProxy updates NPM proxy repository
// endpoint: PUT ​/beta​/repositories​/npm​/proxy​/{repositoryName}
// parameters:
// 		r:
// 			description: NPMProxyRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
//    401: Authentication required
//    403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) UpdateNPMProxy(repositoryName string, r NPMProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/npm/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
