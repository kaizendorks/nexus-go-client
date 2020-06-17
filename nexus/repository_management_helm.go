package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type HelmHostedRepository struct {
	Cleanup *Cleanup `json:"cleanup,omitempty"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type HelmProxyRepository struct {
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

// CreateHelmHosted creates Helm hosted repository
// endpoint: POST ​/beta​/repositories​/helm​/hosted
// parameters:
// 		r:
// 			description: HelmHostedRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateHelmHosted(r HelmHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/helm/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateHelmHosted updates Helm hosted repository
// endpoint: PUT ​/beta​/repositories​/helm​/hosted​/{repositoryName}
// parameters:
// 		repositoryName:
// 			description: Name of the repository to update
// 		r:
// 			description: HelmHostedRepository config
// responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateHelmHosted(repositoryName string, r HelmHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/helm/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateHelmProxy creates Helm proxy repository
// endpoint: POST /beta​/repositories​/helm​/proxy
// parameters:
// 		r:
// 			description: HelmProxyRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateHelmProxy(r HelmProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/helm/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateHelmProxy updates Helm proxy repository
// endpoint: PUT ​/beta​/repositories​/helm​/proxy​/{repositoryName}
// parameters:
// 		repositoryName:
// 			description: Name of the repository to update
// 		r:
// 			description: HelmProxyRepository config
// responses:
//		204: Repository updated
//		401: Authentication required
// 		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateHelmProxy(repositoryName string, r HelmProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/helm/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
