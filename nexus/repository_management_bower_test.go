package nexus_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestRepositoryManagementBowerGroup() {
	hostedRepositoryName := "bower-hosted-test"
	err := suite.client.RepositoryManagement.CreateBowerHosted(nexus.BowerHostedRepository{
		Name: hostedRepositoryName,
		Storage: &nexus.Storage{
			BlobStoreName: "default",
			WritePolicy:   "ALLOW_ONCE",
		},
	})
	assert.NoError(suite.T(), err)

	name := "bower-group-test"
	// Create
	testRepository := nexus.BowerGroupRepository{
		Group: &nexus.Group{
			MemberNames: []string{hostedRepositoryName},
		},
		Name:   name,
		Online: true,
		Storage: &nexus.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "ALLOW_ONCE",
		},
	}

	err = suite.client.RepositoryManagement.CreateBowerGroup(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateBowerGroup(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
	err = suite.client.RepositoryManagement.Delete(hostedRepositoryName)
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementBowerHosted() {
	name := "bower-hosted-test"
	testRepository := nexus.BowerHostedRepository{
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
	err := suite.client.RepositoryManagement.CreateBowerHosted(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateBowerHosted(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementBowerProxy() {
	name := "bower-proxy-test"
	testRepository := nexus.BowerProxyRepository{
		Bower: &nexus.BowerProxyAttributes{
			RewritePackageUrls: false,
		},
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
	err := suite.client.RepositoryManagement.CreateBowerProxy(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateBowerProxy(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}
