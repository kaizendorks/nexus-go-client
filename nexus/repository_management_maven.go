package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreateMavenHosted creates Maven hosted repository
//	endpoint: POST ​/beta​/repositories​/maven​/hosted
//	parameters:
// 		r: MavenHostedRepository config
//	responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateMavenHosted(r MavenHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/maven/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateMavenHosted updates Maven hosted repository
//	endpoint: PUT ​/beta​/repositories​/maven​/hosted​/{repositoryName}
//	parameters:
// 		repositoryName: Name of the repository to update
// 		r: MavenHostedRepository config
//	responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateMavenHosted(repositoryName string, r MavenHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/maven/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// CreateMavenProxy creates Maven proxy repository
//	endpoint: POST /beta​/repositories​/maven​/proxy
//	parameters:
// 		r: MavenProxyRepository config
//	responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateMavenProxy(r MavenProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/maven/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateMavenProxy updates Maven proxy repository
//	endpoint: PUT ​/beta​/repositories​/maven​/proxy​/{repositoryName}
//	parameters:
// 		repositoryName: Name of the repository to update
// 		r: MavenProxyRepository config
//	responses:
//		204: Repository updated
//		401: Authentication required
// 		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateMavenProxy(repositoryName string, r MavenProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/maven/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
