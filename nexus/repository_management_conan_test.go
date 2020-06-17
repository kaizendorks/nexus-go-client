package nexus_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestRepositoryManagementConanProxy() {
	name := "conan-proxy-test"
	testRepository := nexus.ConanProxyRepository{
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
	err := suite.client.RepositoryManagement.CreateConanProxy(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateConanProxy(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}
