package nexus_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestRepositoryManagementDockerGroup() {
	hostedRepositoryName := "docker-hosted-test"
	err := suite.client.RepositoryManagement.CreateDockerHosted(nexus.DockerHostedRepository{
		Name:   hostedRepositoryName,
		Docker: &nexus.DockerAttributes{},
		Storage: &nexus.Storage{
			BlobStoreName: "default",
			WritePolicy:   "ALLOW_ONCE",
		},
	})
	assert.NoError(suite.T(), err)

	name := "docker-group-test"
	// Create
	testRepository := nexus.DockerGroupRepository{
		Group: &nexus.Group{
			MemberNames: []string{hostedRepositoryName},
		},
		Name:   name,
		Docker: &nexus.DockerAttributes{},
		Online: true,
		Storage: &nexus.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "ALLOW_ONCE",
		},
	}

	err = suite.client.RepositoryManagement.CreateDockerGroup(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateDockerGroup(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
	err = suite.client.RepositoryManagement.Delete(hostedRepositoryName)
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementDockerHosted() {
	name := "docker-hosted-test"
	testRepository := nexus.DockerHostedRepository{
		Cleanup: &nexus.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Name:   name,
		Docker: &nexus.DockerAttributes{},
		Online: true,
		Storage: &nexus.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "ALLOW_ONCE",
		},
	}

	// Create
	err := suite.client.RepositoryManagement.CreateDockerHosted(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateDockerHosted(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestRepositoryManagementDockerProxy() {
	name := "docker-proxy-test"
	testRepository := nexus.DockerProxyRepository{
		Docker: &nexus.DockerAttributes{},
		DockerProxy: &nexus.DockerProxyAttributes{
			IndexType: "REGISTRY",
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
	err := suite.client.RepositoryManagement.CreateDockerProxy(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateDockerProxy(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}
