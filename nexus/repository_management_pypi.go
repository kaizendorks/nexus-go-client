package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PyPiGroupRepository struct {
	Group *Group `json:"group"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type PyPiHostedRepository struct {
	Cleanup *Cleanup `json:"cleanup,omitempty"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type PyPiProxyRepository struct {
	Cleanup *Cleanup `json:"cleanup,omitempty"`

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

// CreatePyPiGroup creates PyPi group repository
// endpoint: POST ​/beta​/repositories​/pypi​/group
// parameters:
// 		r:
// 			description: PyPiGroupRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreatePyPiGroup(r PyPiGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/pypi/group")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdatePyPiGroup updates PyPi group repository
// endpoint: PUT ​/beta​/repositories​/pypi​/group​/{repositoryName}
// parameters:
// 		r:
// 			description: PyPiGroupRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
//    401: Authentication required
//    403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) UpdatePyPiGroup(repositoryName string, r PyPiGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/pypi/group/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreatePyPiHosted create PyPi hosted repository
// endpoint: POST ​/beta​/repositories​/pypi​/hosted
// parameters:
// 		r:
// 			description: PyPiHostedRepository config
// responses:
//    201: Repository created
//    401: Authentication required
//    403: Insufficient permissions
func (a RepositoryManagementAPI) CreatePyPiHosted(r PyPiHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/pypi/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdatePyPiHosted updates PyPi hosted repository
// endpoint: PUT ​/beta​/repositories​/pypi​/hosted​/{repositoryName}
// parameters:
// 		r:
// 			description: PyPiHostedRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
//    401: Authentication required
//    403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) UpdatePyPiHosted(repositoryName string, r PyPiHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/pypi/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreatePyPiProxy creates PyPi proxy repository
// endpoint: POST ​/beta​/repositories​/pypi​/proxy
// parameters:
// 		r:
// 			description: PyPiProxyRepository config
// responses:
//    201: Repository created
//    401: Authentication required
//    403: Insufficient permissions
func (a RepositoryManagementAPI) CreatePyPiProxy(r PyPiProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/pypi/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdatePyPiProxy updates PyPi proxy repository
// endpoint: PUT ​/beta​/repositories​/pypi​/proxy​/{repositoryName}
// parameters:
// 		r:
// 			description: PyPiProxyRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
//    401: Authentication required
//    403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) UpdatePyPiProxy(repositoryName string, r PyPiProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/pypi/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
