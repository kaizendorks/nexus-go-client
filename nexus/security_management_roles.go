package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/kaizendorks/nexus-go-client/models"
)

type SecurityManagementRolesAPI api

// List retrieves a list of all the existing roles
//	api endpoint: GET /beta/security/roles
//	parameters:
// 		rf: RoleFilter object consisting of options to filter the results by.
//	responses:
//		200: successful operation returns RoleResponse slice and nil error
//	 	400: The specified source does not exist
//	 	403: Insufficient permissions to read roles
func (a SecurityManagementRolesAPI) List(rf RoleFilter) ([]RoleResponse, error) {
	path := fmt.Sprintf("beta/security/roles")
	if rf.Source != "" {
		path = fmt.Sprintf("%s?source=%s", path, rf.Source)
	}

	rr := []RoleResponse{}

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return rr, err
	}
	err = json.Unmarshal(resp, &rr)
	return rr, err
}

// Create a new role
//	api endpoint: /beta/security/roles
//	parameters:
// 		r: A Role configuration
//	responses:
// 		200: successful operation returns RoleResponse and nil error
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
//	api endpoint: GET /beta/security/roles/{id}?source={source}
//	parameters:
//		id: The id of the role to look for
// 		rf: RoleFilter object consisting of options to filter the results by.
//	responses:
// 		200: successful operation returns RoleResponse and nil error
// 		400: The specified source does not exist
// 		403: Insufficient permissions to read roles
// 		404: Role not found
func (a SecurityManagementRolesAPI) Get(id string, rf RoleFilter) (RoleResponse, error) {
	path := fmt.Sprintf("beta/security/roles/%s", id)
	if rf.Source != "" {
		path = fmt.Sprintf("%s?source=%s", path, rf.Source)
	}

	rr := RoleResponse{}

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return rr, err
	}
	err = json.Unmarshal(resp, &rr)
	return rr, err
}

// Update an existing role using its ID
//	api endpoint: /beta/security/roles/{id}
//	parameters:
// 		id: The id of the role to update
// 		r: A Role configuration struct
//	responses:
// 		403: Insufficient permissions to update role
// 		404: Role not found
func (a SecurityManagementRolesAPI) Update(id string, r Role) error {
	path := fmt.Sprintf("beta/security/roles/%s", id)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// Delete an existing role using its ID
//	api endpoint: /beta/security/roles/{id}
//	parameters:
// 		id: The id of the role to delete
//	responses:
// 		403: Insufficient permissions to delete role
// 		404: Role not found
func (a SecurityManagementRolesAPI) Delete(id string) error {
	path := fmt.Sprintf("beta/security/roles/%s", id)

	_, err := a.client.sendRequest(http.MethodDelete, path, nil, nil)
	return err
}
