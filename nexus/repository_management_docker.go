package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DockerAttributes struct {
	// Whether to force authentication (Docker Bearer Token Realm required if false)
	ForceBasicAuth bool `json:"forceBasicAuth"`

	// Create an HTTP connector at specified port
	HTTPPort int32 `json:"httpPort,omitempty"`

	// Create an HTTPS connector at specified port
	HTTPSPort int32 `json:"httpsPort,omitempty"`

	// Whether to allow clients to use the V1 API to interact with this repository
	V1Enabled bool `json:"v1Enabled"`
}

type DockerGroupRepository struct {
	Docker *DockerAttributes `json:"docker"`
	Group  *Group            `json:"group"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type DockerHostedRepository struct {
	Cleanup *Cleanup          `json:"cleanup,omitempty"`
	Docker  *DockerAttributes `json:"docker"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type DockerProxyAttributes struct {
	// Type of Docker Index
	// Enum: [HUB REGISTRY CUSTOM]
	IndexType string `json:"indexType,omitempty"`

	// Url of Docker Index to use
	IndexURL string `json:"indexUrl,omitempty"`
}

type DockerProxyRepository struct {
	Cleanup     *Cleanup               `json:"cleanup,omitempty"`
	Docker      *DockerAttributes      `json:"docker"`
	DockerProxy *DockerProxyAttributes `json:"dockerProxy"`
	HTTPClient  *HTTPClient            `json:"httpClient"`

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

// CreateDockerGroup creates Docker group repository
// endpoint: POST ​/beta​/repositories​/docker​/group
// parameters:
// 		r:
// 			description: DockerGroupRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateDockerGroup(r DockerGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/docker/group")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateDockerGroup updates Docker group repository
// endpoint: PUT ​/beta​/repositories​/docker​/group​/{repositoryName}
// parameters:
// 		r:
// 			description: DockerGroupRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateDockerGroup(repositoryName string, r DockerGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/docker/group/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateDockerHosted create Docker hosted repository
// endpoint: POST ​/beta​/repositories​/docker​/hosted
// parameters:
// 		r:
// 			description: DockerHostedRepository config
// responses:
//    201: Repository created
//    401: Authentication required
//    403: Insufficient permissions
func (a RepositoryManagementAPI) CreateDockerHosted(r DockerHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/docker/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateDockerHosted updates Docker hosted repository
// endpoint: PUT ​/beta​/repositories​/docker​/hosted​/{repositoryName}
// parameters:
// 		r:
// 			description: DockerHostedRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateDockerHosted(repositoryName string, r DockerHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/docker/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateDockerProxy creates Docker proxy repository
// endpoint: POST ​/beta​/repositories​/docker​/proxy
// parameters:
// 		r:
// 			description: DockerProxyRepository config
// responses:
//    201: Repository created
//    401: Authentication required
//    403: Insufficient permissions
func (a RepositoryManagementAPI) CreateDockerProxy(r DockerProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/docker/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateDockerProxy updates Docker proxy repository
// endpoint: PUT ​/beta​/repositories​/docker​/proxy​/{repositoryName}
// parameters:
// 		r:
// 			description: DockerProxyRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateDockerProxy(repositoryName string, r DockerProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/docker/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
