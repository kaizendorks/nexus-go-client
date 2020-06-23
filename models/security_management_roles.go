package models

type Role struct {
	Description string `json:"description,omitempty"`
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`

	// The list of privileges assigned to this role.
	// Unique: true
	Privileges []string `json:"privileges"`

	// The list of roles assigned to this role.
	// Unique: true
	Roles []string `json:"roles"`
}

type RoleResponse struct {
	Description string   `json:"description,omitempty"`
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Privileges  []string `json:"privileges"`
	Roles       []string `json:"roles"`

	// The user source which is the origin of this role.
	Source string `json:"source,omitempty"`
}

type RoleFilter struct {
	// (Optional) - The ID of the user source to filter the roles by. Can be fetched using SecurityManagementAPI.List().
	Source string
}
