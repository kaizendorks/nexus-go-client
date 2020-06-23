package models

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
