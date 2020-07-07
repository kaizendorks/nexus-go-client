package nexus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	. "github.com/kaizendorks/nexus-go-client/models"
)

type SecurityManagementUsersAPI api

// List retrieves a list of all the existing users. Note if the source is not 'default' the response is limited to 100 users.
//	api endpoint: GET /beta/security/users
//	parameters:
//		uf: UserFilter object consisting of options to filter the results by.
//	responses:
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
//	api endpoint: /beta/security/users
//	parameters:
// 		nu: A NewUser configuration (a representation of the user to create).
//	responses:
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
//	api endpoint: /beta/security/users/{id}
//	parameters:
// 		id: The userid of the user to update
// 		u: A User configuration (A representation of the user to update).
//	responses:
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
//	api endpoint: /beta/security/users/{id}
//	parameters:
// 		id: The userId of the user to delete
//	responses:
// 		403: The user does not have permission to perform the operation.
// 		404: User not found
func (a SecurityManagementUsersAPI) Delete(id string) error {
	path := fmt.Sprintf("beta/security/users/%s", id)

	_, err := a.client.sendRequest(http.MethodDelete, path, nil, nil)
	return err
}

// ChangePassword change an existing user's password.
//	api endpoint: /beta/security/users/{id}/change-password
//	parameters:
// 		id: The userId of the user to delete
// 		newPassword: The new password to use.
//	responses:
// 		400: Password was not supplied in the body of the request
// 		403: The user does not have permission to perform the operation.
// 		404: User not found
func (a SecurityManagementUsersAPI) ChangePassword(id, newPassword string) error {
	path := fmt.Sprintf("beta/security/users/%s/change-password", id)

	headers := map[string]string{
		"Content-Type": "text/plain",
	}

	_, err := a.client.sendRequest(http.MethodPut, path, strings.NewReader(newPassword), headers)
	return err
}
