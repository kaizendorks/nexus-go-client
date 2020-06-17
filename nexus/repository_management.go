package nexus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RepositoryManagementAPI api

type Repository struct {
	// Component format held in this repository
	Format string `json:"format,omitempty"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name,omitempty"`

	// Whether this repository accepts incoming requests
	// Required: true
	Online bool `json:"online"`

	// Controls if deployments of and updates to artifacts are allowed
	// Enum: [hosted proxy group]
	Type string `json:"type,omitempty"`

	// URL to the repository
	URL string `json:"url,omitempty"`
}

type Storage struct {
	// Blob store used to store repository contents
	BlobStoreName string `json:"blobStoreName,omitempty"`

	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation bool `json:"strictContentTypeValidation"`

	// Controls if deployments of and updates to assets are allowed
	// Enum: [ALLOW, DENY, ALLOW_ONCE]
	WritePolicy string `json:"writePolicy"`
}

type Cleanup struct {
	// Components that match any of the applied policies will be deleted
	PolicyNames []string `json:"policyNames"`
}

type Proxy struct {
	// How long to cache artifacts before rechecking the remote repository (in minutes)
	ContentMaxAge int32 `json:"contentMaxAge"`

	// How long to cache metadata before rechecking the remote repository (in minutes)
	MetadataMaxAge int32 `json:"metadataMaxAge"`

	// Location of the remote repository being proxied
	RemoteURL string `json:"remoteUrl,omitempty"`
}

type NegativeCache struct {
	// Whether to cache responses for content not present in the proxied repository
	Enabled bool `json:"enabled"`

	// How long to cache the fact that a file was not found in the repository (in minutes)
	TimeToLive int32 `json:"timeToLive"`
}

type Group struct {
	// Member repositories' names
	MemberNames []string `json:"memberNames"`
}

type HTTPClientConnectionAuthentication struct {
	NTLMDomain string `json:"ntlmDomain,omitempty"`
	NTLMHost   string `json:"ntlmHost,omitempty"`

	// Authentication type
	// Enum: [username ntlm]
	Type string `json:"type,omitempty"`

	Username string `json:"username,omitempty"`
}

type HTTPClientConnection struct {
	// Whether to enable redirects to the same location (may be required by some servers)
	EnableCircularRedirects bool `json:"enableCircularRedirects,omitempty"`

	// Whether to allow cookies to be stored and used
	EnableCookies bool `json:"enableCookies,omitempty"`

	// Total retries if the initial connection attempt suffers a timeout
	// Maximum: 10
	// Minimum: 0
	Retries int32 `json:"retries,omitempty"`

	// Seconds to wait for activity before stopping and retrying the connection
	// Maximum: 3600
	// Minimum: 1
	Timeout int32 `json:"timeout,omitempty"`

	// Custom fragment to append to User-Agent header in HTTP requests
	UserAgentSuffix string `json:"userAgentSuffix,omitempty"`
}

type HTTPClient struct {
	Authentication *HTTPClientConnectionAuthentication `json:"authentication,omitempty"`

	// Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive
	AutoBlock bool `json:"autoBlock"`

	// Whether to block outbound connections on the repository
	Blocked bool `json:"blocked"`

	Connection *HTTPClientConnection `json:"connection,omitempty"`
}

// List returns the list of repositories
// api endpoint: GET /beta​/repositories
// responses:
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
// api endpoint: DELETE /beta​/repositories​/{repositoryName}
// parameters:
// 		repositoryName:
// 			description: Name of the repository to delete
// 			required: true
// responses:
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
// api endpoint: POST ​/beta​/repositories​/{repositoryName}​/health-check
// parameters:
// 		repositoryName:
// 			description: Name of the repository to enable Repository Health Check for
// 			required: true
// responses:
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
// api endpoint: DELETE ​ /beta​/repositories​/{repositoryName}​/health-check
// parameters:
// 		repositoryName:
// 			description: Name of the repository to disable Repository Health Check for
// 			required: true
// responses:
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
// api endpoint: POST /beta​/repositories​/{repositoryName}​/invalidate-cache
// parameters:
// 		repositoryName
// 			description: Name of the repository to invalidate cache
// 			required: true
// responses:
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
// api endpoint: POST ​/beta​/repositories​/{repositoryName}​/rebuild-index
// parameters:
// 		repositoryName
//    	description: Name of the repository to rebuild index
//      required: true
// responses:
// 		204: Repository search index rebuild has been scheduled
//    400: Repository is not of hosted or proxy type
//    401: Authentication required
//    403: Insufficient permissions
//    404: Repository not found
func (a RepositoryManagementAPI) RebuildIndex(repositoryName string) error {
	path := fmt.Sprintf("beta/repositories/%s/rebuild-index", repositoryName)

	_, err := a.client.sendRequest(http.MethodPost, path, nil, nil)
	return err
}
