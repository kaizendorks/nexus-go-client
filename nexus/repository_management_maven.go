package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MavenHostedRepository struct {
	Cleanup *Cleanup         `json:"cleanup,omitempty"`
	Maven   *MavenAttributes `json:"maven"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type MavenAttributes struct {
	// Validate that all paths are maven artifact or metadata paths
	// Enum: [STRICT PERMISSIVE]
	LayoutPolicy string `json:"layoutPolicy,omitempty"`

	// What type of artifacts does this repository store?
	// Enum: [RELEASE SNAPSHOT MIXED]
	VersionPolicy string `json:"versionPolicy,omitempty"`
}

type MavenProxyRepository struct {
	Cleanup    *Cleanup         `json:"cleanup,omitempty"`
	HTTPClient *HTTPClient      `json:"httpClient"`
	Maven      *MavenAttributes `json:"maven"`

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

// CreateMavenHosted creates Maven hosted repository
// endpoint: POST ​/beta​/repositories​/maven​/hosted
// parameters:
// 		r:
// 			description: MavenHostedRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateMavenHosted(r MavenHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/maven/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateMavenHosted updates Maven hosted repository
// endpoint: PUT ​/beta​/repositories​/maven​/hosted​/{repositoryName}
// parameters:
// 		repositoryName:
// 			description: Name of the repository to update
// 		r:
// 			description: MavenHostedRepository config
// responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateMavenHosted(repositoryName string, r MavenHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/maven/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateMavenProxy creates Maven proxy repository
// endpoint: POST /beta​/repositories​/maven​/proxy
// parameters:
// 		r:
// 			description: MavenProxyRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateMavenProxy(r MavenProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/maven/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateMavenProxy updates Maven proxy repository
// endpoint: PUT ​/beta​/repositories​/maven​/proxy​/{repositoryName}
// parameters:
// 		repositoryName:
// 			description: Name of the repository to update
// 		r:
// 			description: MavenProxyRepository config
// responses:
//		204: Repository updated
//		401: Authentication required
// 		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateMavenProxy(repositoryName string, r MavenProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/maven/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
