package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreatePyPiGroup creates PyPi group repository
//	endpoint: POST /beta/repositories/pypi/group
//	parameters:
// 		r: PyPiGroupRepository config
//	responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreatePyPiGroup(r PyPiGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/pypi/group")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdatePyPiGroup updates PyPi group repository
//	endpoint: PUT /beta/repositories/pypi/group/{repositoryName}
//	parameters:
// 		r: PyPiGroupRepository config
// 		repositoryName: Name of the repository to update
//	responses:
// 		204: Repository updated
//		401: Authentication required
//		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdatePyPiGroup(repositoryName string, r PyPiGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/pypi/group/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreatePyPiHosted create PyPi hosted repository
//	endpoint: POST /beta/repositories/pypi/hosted
//	parameters:
// 		r: PyPiHostedRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreatePyPiHosted(r PyPiHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/pypi/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdatePyPiHosted updates PyPi hosted repository
//	endpoint: PUT /beta/repositories/pypi/hosted/{repositoryName}
//	parameters:
// 		r: PyPiHostedRepository config
// 		repositoryName: Name of the repository to update
//	responses:
// 		204: Repository updated
//		401: Authentication required
//		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdatePyPiHosted(repositoryName string, r PyPiHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/pypi/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreatePyPiProxy creates PyPi proxy repository
//	endpoint: POST /beta/repositories/pypi/proxy
//	parameters:
// 		r: PyPiProxyRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreatePyPiProxy(r PyPiProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/pypi/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdatePyPiProxy updates PyPi proxy repository
//	endpoint: PUT /beta/repositories/pypi/proxy/{repositoryName}
//	parameters:
// 		r: PyPiProxyRepository config
// 		repositoryName: Name of the repository to update
//	responses:
// 		204: Repository updated
//		401: Authentication required
//		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdatePyPiProxy(repositoryName string, r PyPiProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/pypi/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
