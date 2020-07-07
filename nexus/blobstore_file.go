package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

// CreateFileStore create a new file blob store
//	endpoint: POST /beta/blobstores/file
//	parameters:
// 		bs: set of config options to use when creating the new blob store
//	responses:
// 		200: Successful operation
func (a BlobStoreAPI) CreateFileStore(bs FileBlobStoreConfig) error {
	path := fmt.Sprintf("beta/blobstores/file")

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(bs)

	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	return err
}

// GetFileStore gets a file blob store configuration by name
//	endpoint: GET /beta/blobstores/file/{name}
//	parameters:
// 		blobStoreName: The name of the file blob store to read
//	responses:
// 		200: successful operation return FileBlobStore and nil error
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
//	endpoint: /beta/blobstores/file/{name}
//	parameters:
// 		blobStoreName: The name of the file blob store to update
// 		bs: FileBlobStore configuration
//	responses:
// 		default: successful operation
func (a BlobStoreAPI) UpdateFileStore(blobStoreName string, bs FileBlobStore) error {
	path := fmt.Sprintf("beta/blobstores/file/%s", blobStoreName)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(bs)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}
