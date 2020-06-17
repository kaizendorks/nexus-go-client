package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type APTHostedRepositoryAttributes struct {
	Distribution string `json:"distribution,omitempty"`
}

type APTSigningRepositoriesAttributes struct {
	// PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)
	Keypair string `json:"keypair,omitempty"`

	// Passphrase to access PGP signing key
	Passphrase string `json:"passphrase,omitempty"`
}

type APTHostedRepository struct {
	APT        *APTHostedRepositoryAttributes    `json:"apt"`
	APTSigning *APTSigningRepositoriesAttributes `json:"aptSigning"`
	Cleanup    *Cleanup                          `json:"cleanup,omitempty"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type APTProxyRepositoriesAttributes struct {
	Distribution string `json:"distribution,omitempty"`
	Flat         bool   `json:"flat"`
}

type APTProxyRepository struct {
	APT        *APTProxyRepositoriesAttributes `json:"apt"`
	Cleanup    *Cleanup                        `json:"cleanup,omitempty"`
	HTTPClient *HTTPClient                     `json:"httpClient"`

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

// CreateAPTHosted creates APT hosted repository
// endpoint: POST ​/beta​/repositories​/apt​/hosted
// parameters:
// 		r:
// 			description: APTHostedRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateAPTHosted(r APTHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/apt/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateAPTHosted updates APT hosted repository
// endpoint: PUT ​/beta​/repositories​/apt​/hosted​/{repositoryName}
// parameters:
// 		repositoryName:
// 			description: Name of the repository to update
// 		r:
// 			description: APTHostedRepository config
// responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateAPTHosted(repositoryName string, r APTHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/apt/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateAPTProxy creates APT proxy repository
// endpoint: POST /beta​/repositories​/apt​/proxy
// parameters:
// 		r:
// 			description: APTProxyRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateAPTProxy(r APTProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/apt/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateAPTProxy updates APT proxy repository
// endpoint: PUT ​/beta​/repositories​/apt​/proxy​/{repositoryName}
// parameters:
// 		repositoryName:
// 			description: Name of the repository to update
// 		r:
// 			description: APTProxyRepository config
// responses:
//		204: Repository updated
//		401: Authentication required
// 		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateAPTProxy(repositoryName string, r APTProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/apt/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
