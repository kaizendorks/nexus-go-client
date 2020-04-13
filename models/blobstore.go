package models

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
