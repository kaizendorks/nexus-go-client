package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type YUMAttributes struct {
	// Validate that all paths are RPMs or yum metadata
	// Enum: [PERMISSIVE STRICT]
	DeployPolicy string `json:"deployPolicy,omitempty"`

	// Specifies the repository depth where repodata folder(s) are created
	RepodataDepth int32 `json:"repodataDepth"`
}

type YUMHostedRepository struct {
	Cleanup *Cleanup `json:"cleanup,omitempty"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage       `json:"storage"`
	YUM     *YUMAttributes `json:"yum"`
}

// CreateYUMHosted create YUM hosted repository
// endpoint: POST ​/beta​/repositories​/yum​/hosted
// parameters:
// 		r:
// 			description: YUMHostedRepository config
// responses:
//    201: Repository created
//    401: Authentication required
//    403: Insufficient permissions
func (a RepositoryManagementAPI) CreateYUMHosted(r YUMHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/yum/hosted")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// UpdateYUMHosted updates YUM hosted repository
// endpoint: PUT ​/beta​/repositories​/yum​/hosted​/{repositoryName}
// parameters:
// 		r:
// 			description: YUMHostedRepository config
// 		repositoryName:
// 			description: Name of the repository to update
// responses:
//  	204: Repository updated
//    401: Authentication required
// 		403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) UpdateYUMHosted(repositoryName string, r YUMHostedRepository) error {
	path := fmt.Sprintf("beta/repositories/yum/hosted/%s", repositoryName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
