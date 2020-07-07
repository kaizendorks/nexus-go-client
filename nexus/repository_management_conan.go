package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreateConanProxy creates new Conan proxy repository
//	endpoint: POST /beta/repositories/conan/proxy
//	parameters:
// 		r: ConanProxyRepository config
//	responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateConanProxy(r ConanProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/conan/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateConanProxy updates Conan proxy repository
//	endpoint: PUT /beta/repositories/conan/proxy/{repositoryName}
//	parameters:
// 		repositoryName: Name of the repository to update
// 		r: ConanProxyRepository config
//	responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateConanProxy(repositoryName string, r ConanProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/conan/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
