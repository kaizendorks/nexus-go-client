package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreateBowerGroup creates Bower group repository
//	endpoint: POST /beta/repositories/bower/group
//	parameters:
// 		r: BowerGroupRepository config
//	responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateBowerGroup(r BowerGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/bower/group")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateBowerGroup updates Bower group repository
//	endpoint: PUT /beta/repositories/bower/group/{repositoryName}
//	parameters:
// 		r: BowerGroupRepository config
// 		repositoryName: Name of the repository to update
//	responses:
// 		204: Repository updated
//		401: Authentication required
//		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateBowerGroup(repositoryName string, r BowerGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/bower/group/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateBowerHosted create Bower hosted repository
//	endpoint: POST /beta/repositories/bower/hosted
//	parameters:
// 		r: BowerHostedRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateBowerHosted(r BowerHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/bower/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateBowerHosted updates Bower hosted repository
//	endpoint: PUT /beta/repositories/bower/hosted/{repositoryName}
//	parameters:
// 		r: BowerHostedRepository config
// 		repositoryName: Name of the repository to update
//	responses:
// 		204: Repository updated
//		401: Authentication required
//		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateBowerHosted(repositoryName string, r BowerHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/bower/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateBowerProxy creates Bower proxy repository
//	endpoint: POST /beta/repositories/bower/proxy
//	parameters:
// 		r: BowerProxyRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateBowerProxy(r BowerProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/bower/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateBowerProxy updates Bower proxy repository
//	endpoint: PUT /beta/repositories/bower/proxy/{repositoryName}
//	parameters:
// 		r: BowerProxyRepository config
// 		repositoryName: Name of the repository to update
//	responses:
// 		204: Repository updated
//		401: Authentication required
//		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateBowerProxy(repositoryName string, r BowerProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/bower/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
