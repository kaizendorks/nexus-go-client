package nexus

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
//
// 	. "github.com/kaizendorks/nexus-go-client/models"
// )

// CreateS3Store creates an S3 blob store
//	api endpoint: POST /beta/blobstores/s3
//	parameters:
// 		bs: S3BlobStore config object.
//	responses:
// 		201: S3 blob store created
//	 	401: Authentication required
//	 	403: Insufficient permissions
// func (a BlobStoreAPI) CreateS3Store(bs S3BlobStore) error {
// 	path := fmt.Sprintf("beta/blobstores/s3")

// 	b := new(bytes.Buffer)
// 	json.NewEncoder(b).Encode(bs)

// 	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
// 	return err
// }

// GetS3Store fetch a S3 blob store configuration
//	api endpoint: GET /beta/blobstores/s3/{name}
//	parameters:
// 		blobStoreName: Name of the blob store configuration to fetch
//	responses:
//	 	200: OK return S3BlobStore and nil error
//	 	400: Specified S3 blob store doesn't exist
//	 	401: Authentication required
//	 	403: Insufficient permissions
// func (a BlobStoreAPI) GetS3Store(blobStoreName string) (S3BlobStore, error) {
// 	bs := S3BlobStore{}
// 	path := fmt.Sprintf("beta/blobstores/s3/%s", blobStoreName)

// 	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
// 	if err != nil {
// 		return bs, err
// 	}

// 	err = json.Unmarshal(resp, &bs)
// 	return bs, err
// }

// UpdateS3Store updates an S3 blob store configuration
//	api endpoint: PUT /beta/blobstores/s3/{name}
//	parameters:
// 		bs: S3BlobStore config object
// 		blobStoreName: Name of the blob store to update
//	responses:
//	 	204: S3 blob store updated
//	 	400: Specified S3 blob store doesn't exist
//	 	401: Authentication required
//	 	403: Insufficient permissions
// func (a BlobStoreAPI) UpdateS3Store(blobStoreName string, bs S3BlobStore) error {
// 	path := fmt.Sprintf("beta/blobstores/s3/%s", blobStoreName)

// 	b := new(bytes.Buffer)
// 	json.NewEncoder(b).Encode(bs)

// 	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
// 	return err
// }
