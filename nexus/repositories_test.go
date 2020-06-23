package nexus_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/models"
	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestRepositoriesList() {
	expected := []models.RepositoryV1{
		models.RepositoryV1{Attributes: map[string]interface{}{}, Format: "maven2", Name: "maven-snapshots", Type: "hosted", URL: "http://nexus:8081/repository/maven-snapshots"},
		models.RepositoryV1{Attributes: map[string]interface{}{"proxy": map[string]interface{}{"remoteUrl": "https://repo1.maven.org/maven2/"}}, Format: "maven2", Name: "maven-central", Type: "proxy", URL: "http://nexus:8081/repository/maven-central"},
		models.RepositoryV1{Attributes: map[string]interface{}{}, Format: "nuget", Name: "nuget-group", Type: "group", URL: "http://nexus:8081/repository/nuget-group"},
		models.RepositoryV1{Attributes: map[string]interface{}{"proxy": map[string]interface{}{"remoteUrl": "https://www.nuget.org/api/v2/"}}, Format: "nuget", Name: "nuget.org-proxy", Type: "proxy", URL: "http://nexus:8081/repository/nuget.org-proxy"},
		models.RepositoryV1{Attributes: map[string]interface{}{}, Format: "maven2", Name: "maven-releases", Type: "hosted", URL: "http://nexus:8081/repository/maven-releases"},
		models.RepositoryV1{Attributes: map[string]interface{}{}, Format: "nuget", Name: "nuget-hosted", Type: "hosted", URL: "http://nexus:8081/repository/nuget-hosted"},
		models.RepositoryV1{Attributes: map[string]interface{}{}, Format: "maven2", Name: "maven-public", Type: "group", URL: "http://nexus:8081/repository/maven-public"},
	}
	actual, err := suite.client.Repositories.List()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *MockedClientSuite) TestRepositoriesListError() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte("Some Server Error"))
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	_, err := mockedClient.Repositories.List()
	assert.Error(suite.T(), err)
}
