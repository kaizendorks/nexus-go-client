package nexus_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/models"
	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestAssetsList() {
	assets, err := suite.client.Assets.List(models.AssetFilter{
		Repository: "maven-central",
	})

	assert.NoError(suite.T(), err)
	assert.ElementsMatch(suite.T(), assets.Items, []models.Asset{})
	assert.Empty(suite.T(), assets.ContinuationToken)

	_, err = suite.client.Assets.List(models.AssetFilter{
		Repository:        "invalid",
		ContinuationToken: "fake",
	})

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "404 Not Found", err.Error())
}

func (suite *MockedClientSuite) TestAssetsGet() {
	expected := &models.Asset{
		Checksum:    map[string]string{"sha1": "c2eabea90b4b10ec5a26de63ea7516f38d805026", "sha256": "8415f3081edf1ad04b0333c18f5a2cc23e647aa67d5958a1b1b613a20943c6c6"},
		DownloadURL: "https://testurl/repository/chips-docker-registry/v2/-/blobs/sha256:8415f3081edf1ad04b0333c18f5a2cc23e647aa67d5958a1b1b613a20943c6c6",
		Format:      "docker",
		ID:          "Y2hpcHMtZG9ja2VyLXJlZ2lzdHJ5OmM3NzAxNjljMGIyZTNlZDg1M2E3NzI5OGNlM2E4Mjc1",
		Path:        "v2/-/blobs/sha256:8415f3081edf1ad04b0333c18f5a2cc23e647aa67d5958a1b1b613a20943c6c6",
		Repository:  "mocked",
	}

	// generate a test server so we can capture and inspect the request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		mockAsset := `{
			"downloadUrl": "https://testurl/repository/chips-docker-registry/v2/-/blobs/sha256:8415f3081edf1ad04b0333c18f5a2cc23e647aa67d5958a1b1b613a20943c6c6",
			"path": "v2/-/blobs/sha256:8415f3081edf1ad04b0333c18f5a2cc23e647aa67d5958a1b1b613a20943c6c6",
			"id": "Y2hpcHMtZG9ja2VyLXJlZ2lzdHJ5OmM3NzAxNjljMGIyZTNlZDg1M2E3NzI5OGNlM2E4Mjc1",
			"repository": "mocked",
			"format": "docker",
			"checksum": {
				"sha1": "c2eabea90b4b10ec5a26de63ea7516f38d805026",
				"sha256": "8415f3081edf1ad04b0333c18f5a2cc23e647aa67d5958a1b1b613a20943c6c6"
			}
		}`
		w.Write([]byte(mockAsset))
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	actual, err := mockedClient.Assets.Get("mockedAsset")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *NexusClientSuite) TestAssetsGetError() {
	_, err := suite.client.Assets.Get("invalid")

	assert.Error(suite.T(), err)
}

func (suite *MockedClientSuite) TestAssetsDelete() {
	// generate a test server so we can capture and inspect the request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	err := mockedClient.Assets.Delete("mockedAsset")

	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestAssetsDeleteError() {
	err := suite.client.Assets.Delete("Y2hpcHMtZG9ja2VyLXJlZ2lzdHJ5OmM3NzAxNjljMGIyZTNlZDg1M2E3NzI5OGNlM2E4Mjc1")

	assert.Error(suite.T(), err)
}
