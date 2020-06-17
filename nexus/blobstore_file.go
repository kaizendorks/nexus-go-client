package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type FileBlobStore struct {
	// The path to the blobstore contents. This can be an absolute path to anywhere on the system nxrm has access to or it can be a path relative to the sonatype-work directory.
	Path string `json:"path"`

	// Settings to control the soft quota
	SoftQuota *SoftQuota `json:"softQuota"`
}

type FileBlobStoreConfig struct {
	Name string `json:"name,omitempty"`

	// The path to the blobstore contents. This can be an absolute path to anywhere on the system nxrm has access to or it can be a path relative to the sonatype-work directory.
	Path string `json:"path,omitempty"`

	// Settings to control the soft quota
	SoftQuota *SoftQuota `json:"softQuota,omitempty"`
}

// CreateFileStore create a new file blob store
// endpoint: POST /beta​/blobstores​/file
// parameters:
// 		bs:
// 			description: set of config options to use when creating the new blob store
// 			required: false
// 			schema: "#/definitions/FileBlobStoreApiCreateRequest"
// responses:
// 	default: Successful operation
func (a BlobStoreAPI) CreateFileStore(bs FileBlobStoreConfig) error {
	path := fmt.Sprintf("beta/blobstores/file")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(bs)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// GetFileStore gets a file blob store configuration by name
// endpoint: GET /beta​/blobstores​/file​/{name}
// parameters:
// 		name
// 			description: The name of the file blob store to read
// 			required: true
// responses:
// 		200: successful operation
// 		schema: "#/definitions/FileBlobStoreApiModel"
func (a BlobStoreAPI) GetFileStore(blobStoreName string) (FileBlobStore, error) {
	bs := FileBlobStore{}
	path := fmt.Sprintf("beta/blobstores/file/%s", blobStoreName)

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return bs, err
	}

	err = json.Unmarshal(resp, &bs)
	return bs, err
}

// UpdateFileStore updates an existing file blob store configuration by name
// endpoint: ​/beta​/blobstores​/file​/{name}
// parameters:
// 		blobStoreName:
// 			description: The name of the file blob store to update
// 			required: true
// 		bs:
// 			description: The name of the file blob store to updates
// 			required: false
// 			schema: "$ref": "#/definitions/FileBlobStoreApiUpdateRequest"
// responses:
// 		default: successful operation
func (a BlobStoreAPI) UpdateFileStore(blobStoreName string, bs FileBlobStore) error {
	path := fmt.Sprintf("beta/blobstores/file/%s", blobStoreName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(bs)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
