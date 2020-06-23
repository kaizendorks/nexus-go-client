package nexus_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/models"
	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestStatusReadable() {
	assert.NoError(suite.T(), suite.client.Status.Status())
}

func (suite *MockedClientSuite) TestStatusReadableError() {
	// generate a test server so we can capture and inspect the request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(503)
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	err := mockedClient.Status.Status()
	assert.Error(suite.T(), err)
}

func (suite *NexusClientSuite) TestStatusCheck() {
	expected := map[string]models.SystemStatus{
		"Available CPUs":            models.SystemStatus{Healthy: true},
		"Blob Stores":               models.SystemStatus{Healthy: true},
		"Default Admin Credentials": models.SystemStatus{Healthy: false},
		"DefaultRoleRealm":          models.SystemStatus{Healthy: true},
		"File Descriptors":          models.SystemStatus{Healthy: true},
		"Lifecycle Phase":           models.SystemStatus{Healthy: true},
		"Read-Only Detector":        models.SystemStatus{Healthy: true},
		"Scheduler":                 models.SystemStatus{Healthy: true},
		"Thread Deadlock Detector":  models.SystemStatus{Healthy: true},
		"Transactions":              models.SystemStatus{Healthy: true},
	}

	actual, err := suite.client.Status.StatusCheck()

	assert.NoError(suite.T(), err)

	// TODO: find better way to test partial complex structs
	for k, v := range actual {
		assert.Equal(suite.T(), expected[k].Healthy, v.Healthy)
	}
}

func (suite *MockedClientSuite) TestStatusCheckError() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(503)
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	actual, err := mockedClient.Status.StatusCheck()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), models.StatusCheckResponse{}, actual)
}

func (suite *NexusClientSuite) TestStatusWritable() {
	assert.NoError(suite.T(), suite.client.Status.StatusWritable())
}

func (suite *MockedClientSuite) TestStatusWritableError() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(503)
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	err := mockedClient.Status.StatusWritable()
	assert.Error(suite.T(), err)
}
