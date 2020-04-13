package nexus_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/models"
)

func (suite *NexusClientSuite) TestRepositoryManagementGolangGroup() {
	proxyRepositoryName := "golang-proxy-test"
	err := suite.client.RepositoryManagement.CreateGolangProxy(models.GolangProxyRepository{
		HTTPClient: &models.HTTPClient{
			AutoBlock: true,
			Blocked:   true,
		},
		Name: proxyRepositoryName,
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
	})

	assert.NoError(suite.T(), err)

	name := "golang-group-test"
	// Create
	testRepository := models.GolangGroupRepository{
		Group: &models.Group{
			MemberNames: []string{proxyRepositoryName},
		},
		Name:   name,
		Online: true,
		Storage: &models.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "ALLOW_ONCE",
		},
	}

	err = suite.client.RepositoryManagement.CreateGolangGroup(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateGolangGroup(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
	err = suite.client.RepositoryManagement.Delete(proxyRepositoryName)
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementGolangProxy() {
	name := "golang-proxy-test"
	testRepository := models.GolangProxyRepository{
		HTTPClient: &models.HTTPClient{
			AutoBlock: true,
			Blocked:   true,
		},
		Name: name,
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
	err := suite.client.RepositoryManagement.CreateGolangProxy(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateGolangProxy(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}
