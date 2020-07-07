package nexus_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/models"
	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestSecurityManagementUsersList() {
	expected := []models.User{
		models.User{
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
	actual, err := suite.client.SecurityManagementUsers.List(models.UserFilter{UserID: "ad"})
	assert.NoError(suite.T(), err)
	assert.ElementsMatch(suite.T(), expected, actual)

	actual, err = suite.client.SecurityManagementUsers.List(models.UserFilter{UserID: "ad", Source: "default"})
	assert.NoError(suite.T(), err)
	assert.ElementsMatch(suite.T(), expected, actual)

	actual, err = suite.client.SecurityManagementUsers.List(models.UserFilter{Source: "default"})
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

	_, err := mockedClient.SecurityManagementUsers.List(models.UserFilter{})
	assert.Error(suite.T(), err)
}

func (suite *NexusClientSuite) TestSecurityManagementUser() {
	id := "test-user"
	nu := models.NewUser{
		EmailAddress: "test@user.com",
		FirstName:    id,
		LastName:     id,
		Password:     "secret",
		Roles:        []string{"nx-admin"},
		Status:       "active",
		UserID:       id,
	}
	expected := models.User{
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

	_, err = suite.client.SecurityManagementUsers.Create(models.NewUser{})
	assert.Error(suite.T(), err)

	// Update
	expected.EmailAddress = "updated@user.com"
	err = suite.client.SecurityManagementUsers.Update(id, expected)
	assert.NoError(suite.T(), err)

	// Change password
	newPassword := "newpass"
	err = suite.client.SecurityManagementUsers.ChangePassword(id, newPassword)
	assert.NoError(suite.T(), err)

	// Login with new password
	newClient := nexus.NewClient(nexus.ClientConfig{
		Host:     suite.client.Config.Host,
		Username: id,
		Password: newPassword,
	})
	err = newClient.Status.StatusWritable()
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.SecurityManagementUsers.Delete(id)
	assert.NoError(suite.T(), err)
}
