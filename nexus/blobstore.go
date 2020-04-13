package nexus

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

type BlobStoreAPI api

// List all blob stores
//	endpoint: GET /beta​/blobstores
//	responses:
// 		200: successful operation returns BlobStore slice and nill error
func (a BlobStoreAPI) List() ([]BlobStore, error) {
	bb := []BlobStore{}
	path := fmt.Sprintf("beta/blobstores")

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return bb, err
	}
	err = json.Unmarshal(resp, &bb)
	return bb, err
}

// Delete a blob store by name
//	endpoint: DELETE ​ /beta​/blobstores​/{name}
//	parameters:
// 		blobStoreName: The name of the blob store to delete
// 	responses:
// 		default: successful operation returns nill error
func (a BlobStoreAPI) Delete(blobStoreName string) error {
	path := fmt.Sprintf("beta/blobstores/%s", blobStoreName)

	_, err := a.client.sendRequest(http.MethodDelete, path, nil, nil)
	return err
}

// GetQuotaStatus using a blob store name, for checking if the blob store has a quota and is in violation of that quota.
//	endpoint: GET /v1​/blobstores​/{id}​/quota-status
//	parameters:
// 		blobStoreName: The name of the blob store for which to return the quota status
//	responses:
// 		200: successful operation returns QuotaStatusResponse and nil error.
func (a BlobStoreAPI) GetQuotaStatus(blobStoreName string) (QuotaStatusResponse, error) {
	qs := QuotaStatusResponse{}
	path := fmt.Sprintf("v1/blobstores/%s/quota-status", blobStoreName)

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return qs, err
	}

	err = json.Unmarshal(resp, &qs)
	return qs, err
}
