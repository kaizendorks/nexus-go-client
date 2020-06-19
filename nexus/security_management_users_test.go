package nexus_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestSecurityManagementUsersList() {
	expected := []nexus.User{
		nexus.User{
			EmailAddress:  "admin@example.org",
			ExternalRoles: []string{},
			FirstName:     "Administrator",
			LastName:      "User",
			ReadOnly:      false,
			Roles:         []string{"nx-admin"},
			Source:        "default",
			Status:        "active",
			UserID:        "admin",
		},
	}
	actual, err := suite.client.SecurityManagementUsers.List(nexus.UserFilter{UserID: "ad"})
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)

	actual, err = suite.client.SecurityManagementUsers.List(nexus.UserFilter{UserID: "ad", Source: "default"})
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)

	actual, err = suite.client.SecurityManagementUsers.List(nexus.UserFilter{Source: "default"})
	assert.NoError(suite.T(), err)
	assert.NotEqual(suite.T(), expected, actual)
	assert.Contains(suite.T(), actual, expected[0])
}

func (suite *MockedClientSuite) TestSecurityManagementUsersListError() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte("Some Server Error"))
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	_, err := mockedClient.SecurityManagementUsers.List(nexus.UserFilter{})
	assert.Error(suite.T(), err)
}

func (suite *NexusClientSuite) TestSecurityManagementUser() {
	id := "test-user"
	nu := nexus.NewUser{
		EmailAddress: "test@user.com",
		FirstName:    id,
		LastName:     id,
		Password:     "secret",
		Roles:        []string{"nx-admin"},
		Status:       "active",
		UserID:       id,
	}
	expected := nexus.User{
		EmailAddress:  "test@user.com",
		ExternalRoles: []string{},
		FirstName:     id,
		LastName:      id,
		ReadOnly:      false,
		Roles:         []string{"nx-admin"},
		Source:        "default",
		Status:        "active",
		UserID:        id,
	}

	// Create
	actual, err := suite.client.SecurityManagementUsers.Create(nu)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)

	_, err = suite.client.SecurityManagementUsers.Create(nexus.NewUser{})
	assert.Error(suite.T(), err)

	// Update
	expected.EmailAddress = "updated@user.com"
	err = suite.client.SecurityManagementUsers.Update(id, expected)
	assert.NoError(suite.T(), err)

	// Change password
	err = suite.client.SecurityManagementUsers.ChangePassword(id, "newpass")
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.SecurityManagementUsers.Delete(id)
	assert.NoError(suite.T(), err)
}
