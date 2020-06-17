package nexus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BlobStoreAPI api

type SoftQuota struct {
	// The limit in MB.
	Limit int64 `json:"limit"`

	// The type to use such as spaceRemainingQuota, or spaceUsedQuota
	Type string `json:"type"`
}

type BlobStore struct {
	AvailableSpaceInBytes int64      `json:"availableSpaceInBytes"`
	BlobCount             int64      `json:"blobCount"`
	Name                  string     `json:"name"`
	SoftQuota             *SoftQuota `json:"softQuota"`
	TotalSizeInBytes      int64      `json:"totalSizeInBytes"`
	Type                  string     `json:"type"`
}

type QuotaStatusResponse struct {
	BlobStoreName string `json:"blobStoreName,omitempty"`
	IsViolation   bool   `json:"isViolation,omitempty"`
	Message       string `json:"message,omitempty"`
}

// List all blob stores
// endpoint: GET /beta​/blobstores
// responses:
// 		200: Successful operation return BlobStore slice and nill error
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
// endpoint: DELETE ​ /beta​/blobstores​/{name}
// parameters:
// 		blobStoreName
// 			description: The name of the blob store to delete
// 			required: true
// 	responses:
// 		default: Successful operation return nill error
func (a BlobStoreAPI) Delete(blobStoreName string) error {
	path := fmt.Sprintf("beta/blobstores/%s", blobStoreName)

	_, err := a.client.sendRequest(http.MethodDelete, path, nil, nil)
	return err
}

// GetQuotaStatus using a blob store name, for checking if the blob store has a quota and is in violation of that quota.
// endpoint: GET /v1​/blobstores​/{id}​/quota-status
// parameters:
// 		blobStoreName
//      description: the name of the blob store for which to return the quota status
// 			required: true
// responses:
// 		200: Successful operation return QuotaStatusResponse and nil error.
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
