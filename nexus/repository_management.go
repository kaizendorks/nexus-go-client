package nexus

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

type RepositoryManagementAPI api

// List returns the list of repositories
//	api endpoint: GET /beta/repositories
//	responses:
// 		200: Return Repository slice and nil error
// 		401: Authentication required
// 		403: Insufficient permissions
func (a RepositoryManagementAPI) List() ([]Repository, error) {
	rr := []Repository{}
	path := fmt.Sprintf("beta/repositories")

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return rr, err
	}
	err = json.Unmarshal(resp, &rr)
	return rr, err
}

// Delete repository of any format
//	api endpoint: DELETE /beta/repositories/{repositoryName}
//	parameters:
// 		repositoryName: Name of the repository to delete
// 			required: true
//	responses:
// 		204: Repository Health Check disabled
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) Delete(repositoryName string) error {
	path := fmt.Sprintf("beta/repositories/%s", repositoryName)

	_, err := a.client.sendRequest(http.MethodDelete, path, nil, nil)
	return err
}

// EnableHealthCheck enables repository health check. Proxy repositories only.
//	api endpoint: POST /beta/repositories/{repositoryName}/health-check
//	parameters:
// 		repositoryName: Name of the repository to enable Repository Health Check for
// 			required: true
//	responses:
// 		204: Repository Health Check disabled
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
// 		409: EULA not accepted or Repository Health Check capability not active
func (a RepositoryManagementAPI) EnableHealthCheck(repositoryName string) error {
	path := fmt.Sprintf("beta/repositories/%s/health-check", repositoryName)

	_, err := a.client.sendRequest(http.MethodPost, path, nil, nil)
	return err
}

// DisableHealthCheck disables repository health check. Proxy repositories only.
//	api endpoint: DELETE  /beta/repositories/{repositoryName}/health-check
//	parameters:
// 		repositoryName: Name of the repository to disable Repository Health Check for
// 			required: true
//	responses:
// 		204: Repository Health Check disabled
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) DisableHealthCheck(repositoryName string) error {
	path := fmt.Sprintf("beta/repositories/%s/health-check", repositoryName)

	_, err := a.client.sendRequest(http.MethodDelete, path, nil, nil)
	return err
}

// InvalidateCache invalidates repository cache. Proxy or group repositories only.
//	api endpoint: POST /beta/repositories/{repositoryName}/invalidate-cache
//	parameters:
// 		repositoryName Name of the repository to invalidate cache
// 			required: true
//	responses:
// 		204: Repository cache invalidated
// 		400: Repository is not of proxy or group type
// 		401: Authentication required
// 		403: Insufficient permissions
// 		404: Repository not found
func (a RepositoryManagementAPI) InvalidateCache(repositoryName string) error {
	path := fmt.Sprintf("beta/repositories/%s/invalidate-cache", repositoryName)

	_, err := a.client.sendRequest(http.MethodPost, path, nil, nil)
	return err
}

// RebuildIndex schedule a 'Repair - Rebuild repository search' Task. Hosted or proxy repositories only.
//	api endpoint: POST /beta/repositories/{repositoryName}/rebuild-index
//	parameters:
// 		repositoryName
//			description: Name of the repository to rebuild index
//			required: true
//	responses:
// 		204: Repository search index rebuild has been scheduled
//		400: Repository is not of hosted or proxy type
//		401: Authentication required
//		403: Insufficient permissions
//		404: Repository not found
func (a RepositoryManagementAPI) RebuildIndex(repositoryName string) error {
	path := fmt.Sprintf("beta/repositories/%s/rebuild-index", repositoryName)

	_, err := a.client.sendRequest(http.MethodPost, path, nil, nil)
	return err
}
