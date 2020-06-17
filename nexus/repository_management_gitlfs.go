package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GitLFSHostedRepository struct {
	Cleanup *Cleanup `json:"cleanup,omitempty"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

// CreateGitLFSHosted create GitLFS hosted repository
// endpoint: POST ​/beta​/repositories​/gitlfs​/hosted
// parameters:
// 		r:
// 			description: GitLFSHostedRepository config
// responses:
//    201: Repository created
//    401: Authentication required
//    403: Insufficient permissions
func (a RepositoryManagementAPI) CreateGitLFSHosted(r GitLFSHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/gitlfs/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateGitLFSHosted updates GitLFS hosted repository
// endpoint: PUT ​/beta​/repositories​/gitlfs​/hosted​/{repositoryName}
// parameters:
// 		r:
// 			description: GitLFSHostedRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
//  	204: Repository updated
//    401: Authentication required
// 		403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) UpdateGitLFSHosted(repositoryName string, r GitLFSHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/gitlfs/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
