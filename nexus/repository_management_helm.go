package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreateHelmHosted creates Helm hosted repository
//	endpoint: POST ​/beta​/repositories​/helm​/hosted
//	parameters:
// 		r: HelmHostedRepository config
//	responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateHelmHosted(r HelmHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/helm/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateHelmHosted updates Helm hosted repository
//	endpoint: PUT ​/beta​/repositories​/helm​/hosted​/{repositoryName}
//	parameters:
// 		repositoryName: Name of the repository to update
// 		r: HelmHostedRepository config
//	responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateHelmHosted(repositoryName string, r HelmHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/helm/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateHelmProxy creates Helm proxy repository
//	endpoint: POST /beta​/repositories​/helm​/proxy
//	parameters:
// 		r: HelmProxyRepository config
//	responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateHelmProxy(r HelmProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/helm/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateHelmProxy updates Helm proxy repository
//	endpoint: PUT ​/beta​/repositories​/helm​/proxy​/{repositoryName}
//	parameters:
// 		repositoryName: Name of the repository to update
// 		r: HelmProxyRepository config
//	responses:
//		204: Repository updated
//		401: Authentication required
// 		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateHelmProxy(repositoryName string, r HelmProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/helm/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
