package models

type DockerAttributes struct {
	// Whether to force authentication (Docker Bearer Token Realm required if false)
	ForceBasicAuth bool `json:"forceBasicAuth"`

	// Create an HTTP connector at specified port
	HTTPPort int32 `json:"httpPort,omitempty"`

	// Create an HTTPS connector at specified port
	HTTPSPort int32 `json:"httpsPort,omitempty"`

	// Whether to allow clients to use the V1 API to interact with this repository
	V1Enabled bool `json:"v1Enabled"`
}

type DockerGroupRepository struct {
	Docker *DockerAttributes `json:"docker"`
	Group  *Group            `json:"group"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type DockerHostedRepository struct {
	Cleanup *Cleanup          `json:"cleanup,omitempty"`
	Docker  *DockerAttributes `json:"docker"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type DockerProxyAttributes struct {
	// Type of Docker Index
	// Enum: [HUB REGISTRY CUSTOM]
	IndexType string `json:"indexType,omitempty"`

	// Url of Docker Index to use
	IndexURL string `json:"indexUrl,omitempty"`
}

type DockerProxyRepository struct {
	Cleanup     *Cleanup               `json:"cleanup,omitempty"`
	Docker      *DockerAttributes      `json:"docker"`
	DockerProxy *DockerProxyAttributes `json:"dockerProxy"`
	HTTPClient  *HTTPClient            `json:"httpClient"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	NegativeCache *NegativeCache `json:"negativeCache"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Proxy       *Proxy   `json:"proxy"`
	RoutingRule string   `json:"routingRule,omitempty"`
	Storage     *Storage `json:"storage"`
}
