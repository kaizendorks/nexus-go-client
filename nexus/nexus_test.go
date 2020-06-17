package nexus_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/nexus"
)

// Ensure user config is applied correctly.
func TestNexusClientConstructor(t *testing.T) {
	testUsername := "username"
	testPassword := "password"
	testHost := "http://nexusurl/"

	client := nexus.NewClient(nexus.ClientConfig{
		Username: testUsername,
		Password: testPassword,
		Host:     testHost,
	})

	assert.Equal(t, client.Config.Username, testUsername)
	assert.Equal(t, client.Config.Password, testPassword)
	assert.Equal(t, client.Config.Host, testHost)
}

func TestInvalidServer(t *testing.T) {
	client := nexus.NewClient(nexus.ClientConfig{
		Host: "http://fakehost",
	})

	err := client.Status.Status()
	assert.Error(t, err)
}

func (suite *MockedClientSuite) TestErrorResponseWithoutBody() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	err := mockedClient.Status.Status()
	assert.Error(suite.T(), err)

	assert.Equal(suite.T(), "500 Internal Server Error", err.Error())
}

func (suite *MockedClientSuite) TestErrorResponseWithBody() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte("Some Server Error"))
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	err := mockedClient.Status.Status()
	assert.Error(suite.T(), err)

	assert.Equal(suite.T(), "500 Internal Server Error: Some Server Error", err.Error())
}
