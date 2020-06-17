package nexus_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestSecurityManagementRolesList() {
	expected := []nexus.RoleResponse{
		nexus.RoleResponse{
			Description: "Administrator Role",
			ID:          "nx-admin",
			Name:        "nx-admin",
			Privileges:  []string{"nx-all"},
			Roles:       []string{},
			Source:      "default",
		},
		nexus.RoleResponse{
			Description: "Anonymous Role",
			ID:          "nx-anonymous",
			Name:        "nx-anonymous",
			Privileges: []string{
				"nx-healthcheck-read",
				"nx-search-read",
				"nx-repository-view-*-*-read",
				"nx-repository-view-*-*-browse",
			},
			Roles:  []string{},
			Source: "default",
		},
	}
	actual, err := suite.client.SecurityManagementRoles.List()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)

	actual, err = suite.client.SecurityManagementRoles.ListFromSource("default")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)

	actual, err = suite.client.SecurityManagementRoles.ListFromSource("fake")
	assert.Error(suite.T(), err)
}

func (suite *NexusClientSuite) TestSecurityManagementRole() {
	id := "test-role"
	r := nexus.Role{
		Description: "Administrator Role",
		ID:          id,
		Name:        id,
		Privileges:  []string{"nx-all"},
		Roles:       []string{},
	}
	expected := nexus.RoleResponse{
		Description: "Administrator Role",
		ID:          id,
		Name:        id,
		Privileges:  []string{"nx-all"},
		Roles:       []string{},
		Source:      "default",
	}

	// Create
	actual, err := suite.client.SecurityManagementRoles.Create(r)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)

	_, err = suite.client.SecurityManagementRoles.Create(nexus.Role{})
	assert.Error(suite.T(), err)

	// Update
	r.Description = "Updated Role"
	expected.Description = r.Description
	err = suite.client.SecurityManagementRoles.Update(id, r)
	assert.NoError(suite.T(), err)

	// Get
	actual, err = suite.client.SecurityManagementRoles.Get(id)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)

	_, err = suite.client.SecurityManagementRoles.Get("fakeid")
	assert.Error(suite.T(), err)

	actual, err = suite.client.SecurityManagementRoles.GetFromSource(id, "default")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)

	// cleanup
	err = suite.client.SecurityManagementRoles.Delete(id)
	assert.NoError(suite.T(), err)
}
