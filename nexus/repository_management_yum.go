package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreateYUMHosted create YUM hosted repository
//	endpoint: POST ​/beta​/repositories​/yum​/hosted
//	parameters:
// 		r: YUMHostedRepository config
//	responses:
//		201: Repository created
//		401: Authentication required
//		403: Insufficient permissions
func (a RepositoryManagementAPI) CreateYUMHosted(r YUMHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/yum/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateYUMHosted updates YUM hosted repository
//	endpoint: PUT ​/beta​/repositories​/yum​/hosted​/{repositoryName}
//	parameters:
// 		r: YUMHostedRepository config
// 		repositoryName: Name of the repository to update
//	responses:
//		204: Repository updated
//		401: Authentication required
// 		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) UpdateYUMHosted(repositoryName string, r YUMHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/yum/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
