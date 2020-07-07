package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreateDockerGroup creates Docker group repository
//	endpoint: POST /beta/repositories/docker/group
//	parameters:
// 		r: DockerGroupRepository config
//	responses:
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
//	endpoint: PUT /beta/repositories/docker/group/{repositoryName}
//	parameters:
// 		r: DockerGroupRepository config
// 		repositoryName: Name of the repository to update
//	responses:
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
//	endpoint: POST /beta/repositories/docker/hosted
//	parameters:
// 		r: DockerHostedRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateDockerHosted(r DockerHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/docker/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateDockerHosted updates Docker hosted repository
//	endpoint: PUT /beta/repositories/docker/hosted/{repositoryName}
//	parameters:
// 		r: DockerHostedRepository config
// 		repositoryName: Name of the repository to update
//	responses:
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
//	endpoint: POST /beta/repositories/docker/proxy
//	parameters:
// 		r: DockerProxyRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateDockerProxy(r DockerProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/docker/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateDockerProxy updates Docker proxy repository
//	endpoint: PUT /beta/repositories/docker/proxy/{repositoryName}
//	parameters:
// 		r: DockerProxyRepository config
// 		repositoryName: Name of the repository to update
//	responses:
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
