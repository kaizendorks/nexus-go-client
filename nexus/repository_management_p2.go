package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreateP2Proxy creates new P2 proxy repository
//	endpoint: POST ​/beta​/repositories​/p2​/proxy
//	parameters:
// 		r: P2ProxyRepository config
//	responses:
// 		201: Repository created
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateP2Proxy(r P2ProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/p2/proxy")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateP2Proxy updates P2 proxy repository
//	endpoint: PUT ​/beta​/repositories​/p2​/proxy​/{repositoryName}
//	parameters:
// 		repositoryName: Name of the repository to update
// 		r: P2ProxyRepository config
//	responses:
// 		204: Repository updated
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) UpdateP2Proxy(repositoryName string, r P2ProxyRepository) error {
	path := fmt.Sprintf("beta/repositories/p2/proxy/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
