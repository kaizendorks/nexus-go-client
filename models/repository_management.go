package models

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
