package nexus_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/models"
	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestSecurityManagementUserSource() {
	expected := []models.UserSource{
		models.UserSource{ID: "default", Name: "NexusAuthenticatingRealm"},
		models.UserSource{ID: "LDAP", Name: "LdapRealm"},
	}
	actual, err := suite.client.SecurityManagement.List()
	assert.NoError(suite.T(), err)
	assert.ElementsMatch(suite.T(), expected, actual)
}

func (suite *MockedClientSuite) TestSecurityManagementUserSourceError() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte("Some Server Error"))
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	_, err := mockedClient.SecurityManagement.List()
	assert.Error(suite.T(), err)
}
