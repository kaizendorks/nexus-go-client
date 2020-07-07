package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreateNugetGroup creates Nuget group repository
//	endpoint: POST /beta/repositories/nuget/group
//	parameters:
// 		r: NugetGroupRepository config
//	responses:
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
//	endpoint: PUT /beta/repositories/nuget/group/{repositoryName}
//	parameters:
// 		r: NugetGroupRepository config
// 		repositoryName: Name of the repository to update
//	responses:
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
//	endpoint: POST /beta/repositories/nuget/hosted
//	parameters:
// 		r: NugetHostedRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateNugetHosted(r NugetHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/nuget/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateNugetHosted updates Nuget hosted repository
//	endpoint: PUT /beta/repositories/nuget/hosted/{repositoryName}
//	parameters:
// 		r: NugetHostedRepository config
// 		repositoryName: Name of the repository to update
//	responses:
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
//	endpoint: POST /beta/repositories/nuget/proxy
//	parameters:
// 		r: NugetProxyRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateNugetProxy(r NugetProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/nuget/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateNugetProxy updates Nuget proxy repository
//	endpoint: PUT /beta/repositories/nuget/proxy/{repositoryName}
//	parameters:
// 		r: NugetProxyRepository config
// 		repositoryName: Name of the repository to update
//	responses:
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
