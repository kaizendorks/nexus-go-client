package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreateGolangGroup creates Golang group repository
//	endpoint: POST ​/beta​/repositories​/go/group
//	parameters:
// 		r: GolangGroupRepository config
//	responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateGolangGroup(r GolangGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/go/group")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateGolangGroup updates Golang group repository
//	endpoint: PUT ​/beta​/repositories​/go/group​/{repositoryName}
//	parameters:
// 		r: GolangGroupRepository config
// 		repositoryName: Name of the repository to update
//	responses:
// 		204: Repository updated
//		401: Authentication required
//		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateGolangGroup(repositoryName string, r GolangGroupRepository) error {
	path := fmt.Sprintf("beta/repositories/go/group/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateGolangProxy creates Golang proxy repository
//	endpoint: POST ​/beta​/repositories​/go/proxy
//	parameters:
// 		r: GolangProxyRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateGolangProxy(r GolangProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/go/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateGolangProxy updates Golang proxy repository
//	endpoint: PUT ​/beta​/repositories​/go/proxy​/{repositoryName}
//	parameters:
// 		r: GolangProxyRepository config
// 		repositoryName: Name of the repository to update
//	responses:
// 		204: Repository updated
//		401: Authentication required
//		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateGolangProxy(repositoryName string, r GolangProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/go/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
