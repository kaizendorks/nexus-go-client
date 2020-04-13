package nexus_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/models"
)

func (suite *NexusClientSuite) TestRepositoryManagementDockerGroup() {
	hostedRepositoryName := "docker-hosted-test"
	err := suite.client.RepositoryManagement.CreateDockerHosted(models.DockerHostedRepository{
		Name:   hostedRepositoryName,
		Docker: &models.DockerAttributes{},
		Storage: &models.Storage{
			BlobStoreName: "default",
			WritePolicy:   "ALLOW_ONCE",
		},
	})
	assert.NoError(suite.T(), err)

	name := "docker-group-test"
	// Create
	testRepository := models.DockerGroupRepository{
		Group: &models.Group{
			MemberNames: []string{hostedRepositoryName},
		},
		Name:   name,
		Docker: &models.DockerAttributes{},
		Online: true,
		Storage: &models.Storage{
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
	testRepository := models.DockerHostedRepository{
		Cleanup: &models.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Name:   name,
		Docker: &models.DockerAttributes{},
		Online: true,
		Storage: &models.Storage{
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
	testRepository := models.DockerProxyRepository{
		Docker: &models.DockerAttributes{},
		DockerProxy: &models.DockerProxyAttributes{
			IndexType: "REGISTRY",
		},
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
	err := suite.client.RepositoryManagement.CreateDockerProxy(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateDockerProxy(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}
