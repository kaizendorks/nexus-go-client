package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type P2ProxyRepository struct {
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

// POST  ​/beta​/repositories​/p2​/proxy                   Create P2 proxy repository
// summary: Create P2 proxy repository
// parameters:
// 		r:
// 			description: P2ProxyRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateP2Proxy(r P2ProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/p2/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// PUT   ​/beta​/repositories​/p2​/proxy​/{repositoryName}  Update P2 proxy repository
// summary: Update P2 proxy repository
// parameters:
// 		repositoryName:
// 			description: Name of the repository to update
// 		r:
// 			description: P2ProxyRepository config
// responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateP2Proxy(repositoryName string, r P2ProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/p2/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
