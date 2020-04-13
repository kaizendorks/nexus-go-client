package models

type User struct {
	EmailAddress string `json:"emailAddress,omitempty"`

	// The roles which the user has been assigned in an external source, e.g. LDAP group. These cannot be changed within Nexus.
	ExternalRoles []string `json:"externalRoles"`

	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`

	// Indicates whether the user's properties could be modified by Nexus. When false only roles are considered during update.
	ReadOnly bool `json:"readOnly,omitempty"`

	// The roles which the user has been assigned within Nexus.
	Roles []string `json:"roles"`

	// The user source which is the origin of this user. This value cannot be changed.
	Source string `json:"source,omitempty"`

	// The user's status, e.g. active or disabled.
	// Enum: [active locked disabled changepassword]
	Status string `json:"status"`

	// The userid which is required for login. This value cannot be changed.
	UserID string `json:"userId,omitempty"`
}

type NewUser struct {
	EmailAddress string `json:"emailAddress,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	Password     string `json:"password,omitempty"`

	// The roles which the user has been assigned within Nexus.
	Roles []string `json:"roles"`

	// The user's status, e.g. active or disabled.
	// Enum: [active locked disabled changepassword]
	Status string `json:"status"`

	// The userid which is required for login. This value cannot be changed.
	UserID string `json:"userId,omitempty"`
}

type UserFilter struct {
	// An optional term to search userIds for. Matches userId that starts or is equal to the value passed in.
	UserID string

	// An optional user source to restrict the search to. Matches full string.
	Source string
}
