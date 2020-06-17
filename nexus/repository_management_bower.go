package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type BowerProxyAttributes struct {
	// Whether to force Bower to retrieve packages through this proxy repository
	RewritePackageUrls bool `json:"rewritePackageUrls"`
}

type BowerGroupRepository struct {
	Group *Group `json:"group"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type BowerHostedRepository struct {
	Cleanup *Cleanup `json:"cleanup,omitempty"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type BowerProxyRepository struct {
	Bower      *BowerProxyAttributes `json:"bower,omitempty"`
	Cleanup    *Cleanup              `json:"cleanup,omitempty"`
	HTTPClient *HTTPClient           `json:"httpClient"`

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

// CreateBowerGroup creates Bower group repository
// endpoint: POST ​/beta​/repositories​/bower​/group
// parameters:
// 		r:
// 			description: BowerGroupRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateBowerGroup(r BowerGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/bower/group")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateBowerGroup updates Bower group repository
// endpoint: PUT ​/beta​/repositories​/bower​/group​/{repositoryName}
// parameters:
// 		r:
// 			description: BowerGroupRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
//    401: Authentication required
//    403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) UpdateBowerGroup(repositoryName string, r BowerGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/bower/group/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateBowerHosted create Bower hosted repository
// endpoint: POST ​/beta​/repositories​/bower​/hosted
// parameters:
// 		r:
// 			description: BowerHostedRepository config
// responses:
//    201: Repository created
//    401: Authentication required
//    403: Insufficient permissions
func (a RepositoryManagementAPI) CreateBowerHosted(r BowerHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/bower/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateBowerHosted updates Bower hosted repository
// endpoint: PUT ​/beta​/repositories​/bower​/hosted​/{repositoryName}
// parameters:
// 		r:
// 			description: BowerHostedRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
//    401: Authentication required
//    403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) UpdateBowerHosted(repositoryName string, r BowerHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/bower/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateBowerProxy creates Bower proxy repository
// endpoint: POST ​/beta​/repositories​/bower​/proxy
// parameters:
// 		r:
// 			description: BowerProxyRepository config
// responses:
//    201: Repository created
//    401: Authentication required
//    403: Insufficient permissions
func (a RepositoryManagementAPI) CreateBowerProxy(r BowerProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/bower/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateBowerProxy updates Bower proxy repository
// endpoint: PUT ​/beta​/repositories​/bower​/proxy​/{repositoryName}
// parameters:
// 		r:
// 			description: BowerProxyRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
//    401: Authentication required
//    403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) UpdateBowerProxy(repositoryName string, r BowerProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/bower/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
