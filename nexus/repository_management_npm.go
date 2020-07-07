package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreateNPMGroup creates NPM group repository
//	endpoint: POST /beta/repositories/npm/group
//	parameters:
// 		r: NPMGroupRepository config
//	responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateNPMGroup(r NPMGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/npm/group")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateNPMGroup updates NPM group repository
//	endpoint: PUT /beta/repositories/npm/group/{repositoryName}
//	parameters:
// 		r: NPMGroupRepository config
// 		repositoryName: Name of the repository to update
//	responses:
// 		204: Repository updated
//		401: Authentication required
//		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateNPMGroup(repositoryName string, r NPMGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/npm/group/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateNPMHosted create NPM hosted repository
//	endpoint: POST /beta/repositories/npm/hosted
//	parameters:
// 		r: NPMHostedRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateNPMHosted(r NPMHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/npm/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateNPMHosted updates NPM hosted repository
//	endpoint: PUT /beta/repositories/npm/hosted/{repositoryName}
//	parameters:
// 		r: NPMHostedRepository config
// 		repositoryName: Name of the repository to update
//	responses:
// 		204: Repository updated
//		401: Authentication required
//		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateNPMHosted(repositoryName string, r NPMHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/npm/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateNPMProxy creates NPM proxy repository
//	endpoint: POST /beta/repositories/npm/proxy
//	parameters:
// 		r: NPMProxyRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateNPMProxy(r NPMProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/npm/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateNPMProxy updates NPM proxy repository
//	endpoint: PUT /beta/repositories/npm/proxy/{repositoryName}
//	parameters:
// 		r: NPMProxyRepository config
// 		repositoryName: Name of the repository to update
//	responses:
// 		204: Repository updated
//		401: Authentication required
//		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateNPMProxy(repositoryName string, r NPMProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/npm/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
