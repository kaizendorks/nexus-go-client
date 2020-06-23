package nexus_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/models"
	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestBlobStoreList() {
	bb, err := suite.client.BlobStore.List()

	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), bb, 1)
	assert.Equal(suite.T(), bb[0].Name, "default")
}

func (suite *MockedClientSuite) TestBlobStoreListError() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte("Some Server Error"))
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	_, err := mockedClient.BlobStore.List()

	assert.Error(suite.T(), err)
}

func (suite *NexusClientSuite) TestBlobStoreDeleteError() {
	err := suite.client.BlobStore.Delete("default")

	assert.Error(suite.T(), err)
}

func (suite *NexusClientSuite) TestBlobStoreGetQuotaStatus() {
	expected := models.QuotaStatusResponse{
		BlobStoreName: "default",
		IsViolation:   false,
		Message:       "Blob store default has no quota",
	}

	actual, err := suite.client.BlobStore.GetQuotaStatus("default")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)

	_, err = suite.client.BlobStore.GetQuotaStatus("invalid")

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "404 Not Found", err.Error())
}
