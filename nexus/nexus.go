// Package nexus provides a client constructor for creating and configuring the Nexus API client.
package nexus

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// User provided config options.
type ClientConfig struct {
	Host     string
	Username string
	Password string
}

// Store basic auth info and http settings for authenticating with the Nexus API.
type Client struct {
	client *http.Client
	Config ClientConfig

	Assets                  *AssetsAPI
	BlobStore               *BlobStoreAPI
	Repositories            *RepositoriesAPI
	RepositoryManagement    *RepositoryManagementAPI
	Script                  *ScriptAPI
	SecurityManagement      *SecurityManagementAPI
	SecurityManagementRoles *SecurityManagementRolesAPI
	Status                  *StatusAPI
}

type api struct {
	client *Client
}

type NexusError struct {
	Status     string
	StatusCode int
	APIError   error
}

func (r *NexusError) Error() string {
	if r.APIError.Error() != "" {
		return fmt.Sprintf("%s: %v", r.Status, r.APIError)
	} else {
		return fmt.Sprintf("%s", r.Status)
	}
}

// NewClient constructs a Nexus API client object based on the user suplied config.
// This is useful for stuff like appling HTTP client config globally, for examplpe, allowing insecure https.
func NewClient(config ClientConfig) Client {
	// TODO: Read and configure insecure tls if necessary.
	httpClient := &http.Client{}

	c := Client{
		client: httpClient,
		Config: config,
	}

	c.Assets = &AssetsAPI{client: &c}
	c.BlobStore = &BlobStoreAPI{client: &c}
	c.Repositories = &RepositoriesAPI{client: &c}
	c.RepositoryManagement = &RepositoryManagementAPI{client: &c}
	c.Script = &ScriptAPI{client: &c}
	c.SecurityManagement = &SecurityManagementAPI{client: &c}
	c.SecurityManagementRoles = &SecurityManagementRolesAPI{client: &c}
	c.Status = &StatusAPI{client: &c}

	return c
}

func (c Client) sendRequest(method, path string, body io.Reader, headers map[string]string) ([]byte, error) {
	requestURL := fmt.Sprintf("%s/service/rest/%s", c.Config.Host, path)

	request, err := http.NewRequest(method, requestURL, body)
	if err != nil {
		return nil, err
	}
	for k, v := range c.Config.getRequestHeaders(headers) {
		request.Header.Set(k, v)
	}

	resp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, &NexusError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
			APIError:   errors.New(string(respBody)),
		}
	}

	return respBody, nil
}

func (c ClientConfig) getRequestHeaders(headers map[string]string) map[string]string {
	requestHeaders := make(map[string]string)

	// Default headers
	encodedAuth := []byte(c.Username + ":" + c.Password)
	requestHeaders["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString(encodedAuth)
	requestHeaders["Content-Type"] = "application/json"

	// Add other default headers here.

	// Merge user headers with default headers if needed
	for k, v := range headers {
		requestHeaders[k] = v
	}

	return requestHeaders
}
