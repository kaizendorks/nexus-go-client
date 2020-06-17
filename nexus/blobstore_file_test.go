package nexus_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestBlobStoreCreateFileStore() {
	err := suite.client.BlobStore.CreateFileStore(nexus.FileBlobStoreConfig{
		Name: "testFileBlobstore",
		Path: "testfilestorepath",
	})

	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.BlobStore.Delete("testFileBlobstore")
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestBlobStoreUpdateFileStore() {
	err := suite.client.BlobStore.CreateFileStore(nexus.FileBlobStoreConfig{
		Name: "testFileBlobstore",
		Path: "testfilestorepath",
	})

	assert.NoError(suite.T(), err)

	err = suite.client.BlobStore.UpdateFileStore("testFileBlobstore", nexus.FileBlobStore{
		Path: "testfilestorepath2",
	})

	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.BlobStore.Delete("testFileBlobstore")
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestBlobStoreGetFileStore() {
	expected := nexus.FileBlobStore{Path: "default"}

	actual, err := suite.client.BlobStore.GetFileStore("default")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)

	_, err = suite.client.BlobStore.GetFileStore("invalid")
	assert.Error(suite.T(), err)
}
