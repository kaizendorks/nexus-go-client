package nexus_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestRepositoryManagementHelmHosted() {
	name := "helm-hosted-test"
	testRepository := nexus.HelmHostedRepository{
		Cleanup: &nexus.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Name:   name,
		Online: true,
		Storage: &nexus.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "ALLOW_ONCE",
		},
	}

	// Create
	err := suite.client.RepositoryManagement.CreateHelmHosted(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateHelmHosted(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementHelmProxy() {
	name := "helm-proxy-test"
	testRepository := nexus.HelmProxyRepository{
		HTTPClient: &nexus.HTTPClient{
			AutoBlock: true,
			Blocked:   true,
		},
		Name: name,
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
	err := suite.client.RepositoryManagement.CreateHelmProxy(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateHelmProxy(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}
