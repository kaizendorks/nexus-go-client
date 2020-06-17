package nexus

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

type S3Bucket struct {
	// How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable)
	// Required: true
	Expiration *int32 `json:"expiration"`

	// The name of the S3 bucket
	// Required: true
	Name *string `json:"name"`

	// The S3 blob store (i.e S3 object) key prefix
	Prefix string `json:"prefix,omitempty"`

	// The AWS region to create a new S3 bucket in or an existing S3 bucket's region
	// Required: true
	Region *string `json:"region"`
}

type S3Encryption struct {
	EncryptionKey string `json:"encryptionKey,omitempty"`

	// The type of S3 server side encryption to use.
	// Enum: [s3ManagedEncryption kmsManagedEncryption]
	EncryptionType string `json:"encryptionType,omitempty"`
}

type S3BucketSecurity struct {
	// An IAM access key ID for granting access to the S3 bucket
	AccessKeyID string `json:"accessKeyId,omitempty"`

	// An IAM role to assume in order to access the S3 bucket
	Role string `json:"role,omitempty"`

	// The secret access key associated with the specified IAM access key ID
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// An AWS STS session token associated with temporary security credentials which grant access to the S3 bucket
	SessionToken string `json:"sessionToken,omitempty"`
}

type S3AdvancedBucketConnection struct {

	// A custom endpoint URL for third party object stores using the S3 API.
	Endpoint string `json:"endpoint,omitempty"`

	// Setting this flag will result in path-style access being used for all requests.
	ForcePathStyle bool `json:"forcePathStyle,omitempty"`

	// An API signature version which may be required for third party object stores using the S3 API.
	SignerType string `json:"signerType,omitempty"`
}

type S3BucketConfiguration struct {

	// A custom endpoint URL, signer type and whether path style access is enabled
	AdvancedBucketConnection *S3AdvancedBucketConnection `json:"advancedBucketConnection,omitempty"`

	// Details of the S3 bucket such as name and region
	// Required: true
	// Read Only: true
	Bucket *S3Bucket `json:"bucket"`

	// Security details for granting access the S3 API
	BucketSecurity *S3BucketSecurity `json:"bucketSecurity,omitempty"`

	// The type of encryption to use if any
	// Read Only: true
	Encryption *S3Encryption `json:"encryption,omitempty"`
}

type S3BlobStore struct {
	// The S3 specific configuration details for the S3 object that'll contain the blob store.
	// Required: true
	BucketConfiguration *S3BucketConfiguration `json:"bucketConfiguration"`

	// The name of the S3 blob store.
	// Required: true
	Name *string `json:"name"`

	// Settings to control the soft quota.
	SoftQuota *SoftQuota `json:"softQuota,omitempty"`
}

// CreateS3Store creates an S3 blob store
// api endpoint: POST ​/beta​/blobstores​/s3
// parameters:
// 		bs: S3BlobStore config object.
// responses:
// 		201: S3 blob store created
//   	401: Authentication required
//   	403: Insufficient permissions
// func (a BlobStoreAPI) CreateS3Store(bs S3BlobStore) error {
// 	path := fmt.Sprintf("beta/blobstores/s3")

// 	b := new(bytes.Buffer)
// 	json.NewEncoder(b).Encode(bs)

// 	_, err := a.client.sendRequest(http.MethodPost, path, b, nil)
// 	return err
// }

// GetS3Store fetch a S3 blob store configuration
// api endpoint: GET ​/beta​/blobstores​/s3​/{name}
// parameters:
// 		blobStoreName: Name of the blob store configuration to fetch
// responses:
//   	200: OK return S3BlobStore and nil error
//   	400: Specified S3 blob store doesn't exist
//   	401: Authentication required
//   	403: Insufficient permissions
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
// api endpoint: PUT ​/beta​/blobstores​/s3​/{name}
// parameters:
// 		bs: S3BlobStore config object
// 		blobStoreName: Name of the blob store to update
// responses:
//   	204: S3 blob store updated
//   	400: Specified S3 blob store doesn't exist
//   	401: Authentication required
//   	403: Insufficient permissions
// func (a BlobStoreAPI) UpdateS3Store(blobStoreName string, bs S3BlobStore) error {
// 	path := fmt.Sprintf("beta/blobstores/s3/%s", blobStoreName)

// 	b := new(bytes.Buffer)
// 	json.NewEncoder(b).Encode(bs)

// 	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
// 	return err
// }
