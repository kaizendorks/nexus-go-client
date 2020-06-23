package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreateAPTHosted creates APT hosted repository
//	endpoint: POST ​/beta​/repositories​/apt​/hosted
//	parameters:
// 		r: APTHostedRepository config
//	responses:
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
//	endpoint: PUT ​/beta​/repositories​/apt​/hosted​/{repositoryName}
//	parameters:
// 		repositoryName: Name of the repository to update
// 		r: APTHostedRepository config
//	responses:
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
//	endpoint: POST /beta​/repositories​/apt​/proxy
//	parameters:
// 		r: APTProxyRepository config
//	responses:
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
//	endpoint: PUT ​/beta​/repositories​/apt​/proxy​/{repositoryName}
//	parameters:
// 		repositoryName: Name of the repository to update
// 		r: APTProxyRepository config
//	responses:
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
