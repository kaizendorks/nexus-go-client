package nexus_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/models"
)

func (suite *NexusClientSuite) TestRepositoryManagementMavenHosted() {
	name := "maven-hosted-test"
	testRepository := models.MavenHostedRepository{
		Cleanup: &models.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Name: name,
		Maven: &models.MavenAttributes{
			LayoutPolicy:  "PERMISSIVE",
			VersionPolicy: "RELEASE",
		},
		Online: true,
		Storage: &models.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "ALLOW_ONCE",
		},
	}

	// Create
	err := suite.client.RepositoryManagement.CreateMavenHosted(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateMavenHosted(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}

type MavenAttributes struct {
	// Validate that all paths are maven artifact or metadata paths
	// Enum: [STRICT PERMISSIVE]
	LayoutPolicy string `json:"layoutPolicy,omitempty"`

	// What type of artifacts does this repository store?
	// Enum: [RELEASE SNAPSHOT MIXED]
	VersionPolicy string `json:"versionPolicy,omitempty"`
}

func (suite *NexusClientSuite) TestRepositoryManagementMavenProxy() {
	name := "maven-proxy-test"
	testRepository := models.MavenProxyRepository{
		HTTPClient: &models.HTTPClient{
			AutoBlock: true,
			Blocked:   true,
		},
		Name: name,
		Maven: &models.MavenAttributes{
			LayoutPolicy:  "PERMISSIVE",
			VersionPolicy: "RELEASE",
		},
		NegativeCache: &models.NegativeCache{
			Enabled:    true,
			TimeToLive: 1440,
		},
		Online: true,
		Proxy: &models.Proxy{
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
			RemoteURL:      "http://test",
		},
		RoutingRule: "test",
		Storage: &models.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
	}

	// Create
	err := suite.client.RepositoryManagement.CreateMavenProxy(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateMavenProxy(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}
