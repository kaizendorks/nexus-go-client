package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type SecurityManagementUsersAPI api

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

// List retrieves a list of all the existing users. Note if the source is not 'default' the response is limited to 100 users.
// api endpoint: GET ​/beta/security/users
// parameters:
//		uf: UserFilter object consisting of options to filter the results by.
// responses:
// 		200: successful operation returns User slice and nil error
// 		403: The user does not have permission to perform the operation.
func (a SecurityManagementUsersAPI) List(uf UserFilter) ([]User, error) {
	path := fmt.Sprintf("beta/security/users?userId=%s", uf.UserID)
	if uf.Source != "" {
		path = fmt.Sprintf("%s&source=%s", path, uf.Source)
	}
	users := []User{}

	resp, err := a.client.sendRequest(http.MethodGet, path, nil, nil)
	if err != nil {
		return users, err
	}
	err = json.Unmarshal(resp, &users)
	return users, err
}

// Create a new user in the default source.
// api endpoint: /beta/security/users
// parameters:
// 		nu: A NewUser configuration (a representation of the user to create).
// responses:
// 		200: successful operation returns User and nil error
// 		400: Password was not supplied in the body of the request
// 		403: The user does not have permission to perform the operation.
func (a SecurityManagementUsersAPI) Create(nu NewUser) (User, error) {
	path := fmt.Sprintf("beta/security/users")
	u := User{}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(nu)

	resp, err := a.client.sendRequest(http.MethodPost, path, b, nil)
	if err != nil {
		return u, err
	}

	err = json.Unmarshal(resp, &u)
	return u, err
}

// Update an existing user using its ID
// api endpoint: /beta/security/users/{id}
// parameters:
// 		id: The userid of the user to update
// 		u: A User configuration (A representation of the user to update).
// responses:
// 		200: successful operation returns User and nil error
// 		400: Password was not supplied in the body of the request
// 		403: The user does not have permission to perform the operation.
// 		404: User not found
func (a SecurityManagementUsersAPI) Update(id string, u User) error {
	path := fmt.Sprintf("beta/security/users/%s", id)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)

	_, err := a.client.sendRequest(http.MethodPut, path, b, nil)
	return err
}

// Delete an existing user using its ID
// api endpoint: /beta/security/users/{id}
// parameters:
// 		id: The userId of the user to delete
// responses:
// 		403: The user does not have permission to perform the operation.
// 		404: User not found
func (a SecurityManagementUsersAPI) Delete(id string) error {
	path := fmt.Sprintf("beta/security/users/%s", id)

	_, err := a.client.sendRequest(http.MethodDelete, path, nil, nil)
	return err
}

// put:
//     summary: Change a user's password.
//     parameters:
//     - userId
//       description: The userid the request should apply to.
//       required: true
//       type: string
//     - body
//       description: The new password to use.
//       required: false
//       type: string
//     responses:
//       400: Password was not supplied in the body of the request
//       403: The user does not have permission to perform the operation.
//       404: User not found in the system.

// ChangePassword change an existing user's password.
// api endpoint: /beta/security/users/{id}/change-password
// parameters:
// 		id: The userId of the user to delete
// 		p: The new password to use.
// responses:
// 		400: Password was not supplied in the body of the request
// 		403: The user does not have permission to perform the operation.
// 		404: User not found
func (a SecurityManagementUsersAPI) ChangePassword(id, p string) error {
	path := fmt.Sprintf("beta/security/users/%s/change-password", id)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(p)

	headers := map[string]string{
		"Content-Type": "text/plain",
	}

	_, err := a.client.sendRequest(http.MethodPut, path, b, headers)
	return err
}
