package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SecurityManagementRolesAPI api

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

// List retrieves a list of all the existing roles
// api endpoint: GET ​/beta/security/roles
// responses:
// 		200: successful operation return RoleResponse slice and nil error
// 		403: Insufficient permissions to read roles
func (a SecurityManagementRolesAPI) List() ([]RoleResponse, error) {
	path := fmt.Sprintf("beta/security/roles")
	return a.list(path)
}

// ListFromSource retrieves a of roles belonging to a particular user source
// api endpoint: GET ​/beta/security/roles?source={source}
// parameters:
// 		source: The ID of the user source to filter the roles by.
// responses:
//  	200: successful operation return RoleResponse slice and nil error
//   	400: The specified source does not exist
//   	403: Insufficient permissions to read roles
func (a SecurityManagementRolesAPI) ListFromSource(source string) ([]RoleResponse, error) {
	path := fmt.Sprintf("beta/security/roles?source=%s", source)
	return a.list(path)
}

func (a SecurityManagementRolesAPI) list(path string) ([]RoleResponse, error) {
	rr := []RoleResponse{}

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return rr, err
	}
	err = json.Unmarshal(resp, &rr)
	return rr, err
}

// Create a new role
// api endpoint: /beta/security/roles
// parameters:
// 		r: A Role configuration
// responses:
// 		200: successful operation return RoleResponse and nil error
// 		403: Insufficient permissions to create role
func (a SecurityManagementRolesAPI) Create(r Role) (RoleResponse, error) {
	path := fmt.Sprintf("beta/security/roles")
	roleResponse := RoleResponse{}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	resp, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	if err != nil {
		return roleResponse, err
	}

	err = json.Unmarshal(resp, &roleResponse)
	return roleResponse, err
}

// Get retrieves an existing role by ID
// api endpoint: GET ​/beta/security/{id}
// responses:
// 		200: successful operation return RoleResponse and nil error
//   	403: Insufficient permissions to read roles
// 		404: Role not found
func (a SecurityManagementRolesAPI) Get(id string) (RoleResponse, error) {
	path := fmt.Sprintf("beta/security/roles/%s", id)
	return a.get(path)
}

// GetFromSource retrieves an existing role belonging to a particular user source
// api endpoint: GET ​/beta/security/roles/{id}?source={source}
// parameters:
// 		source: The ID of the user source to filter the roles by. Can be fetched using SecurityManagementAPI.List().
// responses:
// 		200: successful operation return RoleResponse and nil error
// 		400: The specified source does not exist
// 		403: Insufficient permissions to read roles
// 		404: Role not found
func (a SecurityManagementRolesAPI) GetFromSource(id, source string) (RoleResponse, error) {
	path := fmt.Sprintf("beta/security/roles/%s?source=%s", id, source)
	return a.get(path)
}

func (a SecurityManagementRolesAPI) get(path string) (RoleResponse, error) {
	rr := RoleResponse{}

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return rr, err
	}
	err = json.Unmarshal(resp, &rr)
	return rr, err
}

// Update an existing role using it's ID
// api endpoint: /beta/security/roles/{id}
// parameters:
// 		id: The id of the role to update
// 		r: A Role configuration struct
// responses:
// 		403: Insufficient permissions to update role
// 		404: Role not found
func (a SecurityManagementRolesAPI) Update(id string, r Role) error {
	path := fmt.Sprintf("beta/security/roles/%s", id)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// Delete an existing role using it's ID
// api endpoint: /beta/security/roles/{id}
// parameters:
// 		id: The id of the role to delete
// responses:
// 		403: Insufficient permissions to delete role
// 		404: Role not found
func (a SecurityManagementRolesAPI) Delete(id string) error {
	path := fmt.Sprintf("beta/security/roles/%s", id)

	_, err := a.client.sendRequest(http.MethodDelete, path, nil, nil)
	return err
}
