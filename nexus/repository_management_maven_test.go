package nexus_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestRepositoryManagementMavenHosted() {
	name := "maven-hosted-test"
	testRepository := nexus.MavenHostedRepository{
		Cleanup: &nexus.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Name: name,
		Maven: &nexus.MavenAttributes{
			LayoutPolicy:  "PERMISSIVE",
			VersionPolicy: "RELEASE",
		},
		Online: true,
		Storage: &nexus.Storage{
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
	testRepository := nexus.MavenProxyRepository{
		HTTPClient: &nexus.HTTPClient{
			AutoBlock: true,
			Blocked:   true,
		},
		Name: name,
		Maven: &nexus.MavenAttributes{
			LayoutPolicy:  "PERMISSIVE",
			VersionPolicy: "RELEASE",
		},
		NegativeCache: &nexus.NegativeCache{
			Enabled:    true,
			TimeToLive: 1440,
		},
		Online: true,
		Proxy: &nexus.Proxy{
			ContentMaxAge:  1440,
			MetadataMaxAge: 1440,
			RemoteURL:      "http://test",
		},
		RoutingRule: "test",
		Storage: &nexus.Storage{
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
