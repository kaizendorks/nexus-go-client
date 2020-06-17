package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type NugetGroupRepository struct {
	Group *Group `json:"group"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type NugetHostedRepository struct {
	Cleanup *Cleanup `json:"cleanup,omitempty"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type NugetProxyAttributes struct {
	// What type of artifacts does this repository store?
	QueryCacheItemMaxAge int32 `json:"queryCacheItemMaxAge,omitempty"`
}

type NugetProxyRepository struct {
	Cleanup    *Cleanup    `json:"cleanup,omitempty"`
	HTTPClient *HTTPClient `json:"httpClient"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	NegativeCache *NegativeCache        `json:"negativeCache"`
	NugetProxy    *NugetProxyAttributes `json:"nugetProxy"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Proxy       *Proxy   `json:"proxy"`
	RoutingRule string   `json:"routingRule,omitempty"`
	Storage     *Storage `json:"storage"`
}

// CreateNugetGroup creates Nuget group repository
// endpoint: POST ​/beta​/repositories​/nuget​/group
// parameters:
// 		r:
// 			description: NugetGroupRepository config
// responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateNugetGroup(r NugetGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/nuget/group")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateNugetGroup updates Nuget group repository
// endpoint: PUT ​/beta​/repositories​/nuget​/group​/{repositoryName}
// parameters:
// 		r:
// 			description: NugetGroupRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateNugetGroup(repositoryName string, r NugetGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/nuget/group/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateNugetHosted create Nuget hosted repository
// endpoint: POST ​/beta​/repositories​/nuget​/hosted
// parameters:
// 		r:
// 			description: NugetHostedRepository config
// responses:
//    201: Repository created
//    401: Authentication required
//    403: Insufficient permissions
func (a RepositoryManagementAPI) CreateNugetHosted(r NugetHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/nuget/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateNugetHosted updates Nuget hosted repository
// endpoint: PUT ​/beta​/repositories​/nuget​/hosted​/{repositoryName}
// parameters:
// 		r:
// 			description: NugetHostedRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateNugetHosted(repositoryName string, r NugetHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/nuget/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateNugetProxy creates Nuget proxy repository
// endpoint: POST ​/beta​/repositories​/nuget​/proxy
// parameters:
// 		r:
// 			description: NugetProxyRepository config
// responses:
//    201: Repository created
//    401: Authentication required
//    403: Insufficient permissions
func (a RepositoryManagementAPI) CreateNugetProxy(r NugetProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/nuget/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateNugetProxy updates Nuget proxy repository
// endpoint: PUT ​/beta​/repositories​/nuget​/proxy​/{repositoryName}
// parameters:
// 		r:
// 			description: NugetProxyRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateNugetProxy(repositoryName string, r NugetProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/nuget/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
