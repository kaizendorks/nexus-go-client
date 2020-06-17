package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ConanProxyRepository struct {
	Cleanup    *Cleanup    `json:"cleanup,omitempty"`
	HTTPClient *HTTPClient `json:"httpClient"`

	// A unique identifier for this repository
	// Required: true
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	NegativeCache *NegativeCache `json:"negativeCache"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Proxy       *Proxy   `json:"proxy"`
	RoutingRule string   `json:"routingRule,omitempty"`
	Storage     *Storage `json:"storage"`
}

// CreateConanProxy creates new Conan proxy repository
// endpoint: POST ​/beta​/repositories​/conan​/proxy
// parameters:
// 		r:
// 			description: ConanProxyRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateConanProxy(r ConanProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/conan/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateConanProxy updates Conan proxy repository
// endpoint: PUT /beta​/repositories​/conan​/proxy​/{repositoryName}
// parameters:
// 		repositoryName:
// 			description: Name of the repository to update
// 		r:
// 			description: ConanProxyRepository config
// responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateConanProxy(repositoryName string, r ConanProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/conan/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
