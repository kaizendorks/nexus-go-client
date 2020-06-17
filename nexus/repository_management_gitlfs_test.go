package nexus_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestRepositoryManagementGitLFSHosted() {
	name := "gitlfs-hosted-test"
	testRepository := nexus.GitLFSHostedRepository{
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
	err := suite.client.RepositoryManagement.CreateGitLFSHosted(testRepository)
	assert.NoError(suite.T(), err)

	// Update
	err = suite.client.RepositoryManagement.UpdateGitLFSHosted(name, testRepository)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.RepositoryManagement.Delete(name)
	assert.NoError(suite.T(), err)
}
