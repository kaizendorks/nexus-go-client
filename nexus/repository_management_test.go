package nexus_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/models"
	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestRepositoryManagementList() {
	expected := []models.Repository{
		models.Repository{Format: "maven2", Name: "maven-snapshots", Online: true, Type: "hosted", URL: "http://nexus:8081/repository/maven-snapshots"},
		models.Repository{Format: "maven2", Name: "maven-central", Online: true, Type: "proxy", URL: "http://nexus:8081/repository/maven-central"},
		models.Repository{Format: "nuget", Name: "nuget-group", Online: true, Type: "group", URL: "http://nexus:8081/repository/nuget-group"},
		models.Repository{Format: "nuget", Name: "nuget.org-proxy", Online: true, Type: "proxy", URL: "http://nexus:8081/repository/nuget.org-proxy"},
		models.Repository{Format: "maven2", Name: "maven-releases", Online: true, Type: "hosted", URL: "http://nexus:8081/repository/maven-releases"},
		models.Repository{Format: "nuget", Name: "nuget-hosted", Online: true, Type: "hosted", URL: "http://nexus:8081/repository/nuget-hosted"},
		models.Repository{Format: "maven2", Name: "maven-public", Online: true, Type: "group", URL: "http://nexus:8081/repository/maven-public"},
	}
	actual, err := suite.client.RepositoryManagement.List()
	assert.NoError(suite.T(), err)
	assert.ElementsMatch(suite.T(), expected, actual)
}

func (suite *MockedClientSuite) TestRepositoryManagementListError() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte("Some Server Error"))
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	_, err := mockedClient.RepositoryManagement.List()
	assert.Error(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementDeleteError() {
	err := suite.client.RepositoryManagement.Delete("fake")
	assert.Error(suite.T(), err)
}

func (suite *MockedClientSuite) TestRepositoryManagementEnableHealthCheck() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	err := mockedClient.RepositoryManagement.EnableHealthCheck("nuget.org-proxy")

	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementEnableHealthCheckError() {
	err := suite.client.RepositoryManagement.EnableHealthCheck("maven-snapshots")
	assert.Error(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementDisableHealthCheck() {
	err := suite.client.RepositoryManagement.DisableHealthCheck("nuget.org-proxy")
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementDisableHealthCheckError() {
	err := suite.client.RepositoryManagement.DisableHealthCheck("maven-snapshots")
	assert.Error(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementInvalidateCache() {
	err := suite.client.RepositoryManagement.InvalidateCache("nuget.org-proxy")
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementInvalidateCacheError() {
	err := suite.client.RepositoryManagement.InvalidateCache("maven-snapshots")
	assert.Error(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementRebuildIndex() {
	err := suite.client.RepositoryManagement.RebuildIndex("nuget.org-proxy")
	assert.NoError(suite.T(), err)
}
