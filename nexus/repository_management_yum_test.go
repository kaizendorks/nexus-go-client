package nexus_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestRepositoryManagementYUMHosted() {
	name := "yum-hosted-test"
	testRepository := nexus.YUMHostedRepository{
		Cleanup: &nexus.Cleanup{
			PolicyNames: []string{"weekly-cleanup"},
		},
		Name: name,
		YUM: &nexus.YUMAttributes{
			DeployPolicy:  "PERMISSIVE",
			RepodataDepth: 2,
		},
		Online: true,
		Storage: &nexus.Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
			WritePolicy:                 "ALLOW_ONCE",
		},
	}

	// Create
	err := suite.client.RepositoryManagement.CreateYUMHosted(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateYUMHosted(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}
