package models

type MavenHostedRepository struct {
	Cleanup *Cleanup         `json:"cleanup,omitempty"`
	Maven   *MavenAttributes `json:"maven"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type MavenAttributes struct {
	// Validate that all paths are maven artifact or metadata paths
	// Enum: [STRICT PERMISSIVE]
	LayoutPolicy string `json:"layoutPolicy,omitempty"`

	// What type of artifacts does this repository store?
	// Enum: [RELEASE SNAPSHOT MIXED]
	VersionPolicy string `json:"versionPolicy,omitempty"`
}

type MavenProxyRepository struct {
	Cleanup    *Cleanup         `json:"cleanup,omitempty"`
	HTTPClient *HTTPClient      `json:"httpClient"`
	Maven      *MavenAttributes `json:"maven"`

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
