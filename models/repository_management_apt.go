package models

type APTHostedRepositoryAttributes struct {
	Distribution string `json:"distribution,omitempty"`
}

type APTSigningRepositoriesAttributes struct {
	// PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)
	Keypair string `json:"keypair,omitempty"`

	// Passphrase to access PGP signing key
	Passphrase string `json:"passphrase,omitempty"`
}

type APTHostedRepository struct {
	APT        *APTHostedRepositoryAttributes    `json:"apt"`
	APTSigning *APTSigningRepositoriesAttributes `json:"aptSigning"`
	Cleanup    *Cleanup                          `json:"cleanup,omitempty"`

	// A unique identifier for this repository
	// Pattern: ^[a-zA-Z0-9\-]{1}[a-zA-Z0-9_\-\.]*$
	Name string `json:"name"`

	// Whether this repository accepts incoming requests
	Online bool `json:"online"`

	Storage *Storage `json:"storage"`
}

type APTProxyRepositoriesAttributes struct {
	Distribution string `json:"distribution,omitempty"`
	Flat         bool   `json:"flat"`
}

type APTProxyRepository struct {
	APT        *APTProxyRepositoriesAttributes `json:"apt"`
	Cleanup    *Cleanup                        `json:"cleanup,omitempty"`
	HTTPClient *HTTPClient                     `json:"httpClient"`

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
