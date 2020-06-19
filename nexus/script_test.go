package nexus_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/kaizendorks/nexus-go-client/nexus"
)

func (suite *NexusClientSuite) TestScriptList() {
	ss, err := suite.client.Script.List()

	assert.NoError(suite.T(), err)
	assert.ElementsMatch(suite.T(), ss, []nexus.Script{})
}

func (suite *MockedClientSuite) TestScriptListError() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte("Some Server Error"))
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	_, err := mockedClient.Script.List()
	assert.Error(suite.T(), err)
}

func (suite *NexusClientSuite) TestScriptWithoutParams() {
	name := "test-script"
	testScript := nexus.Script{
		Name:    name,
		Content: "return 'Hello World!'",
		Type:    "groovy",
	}

	// Create
	err := suite.client.Script.Create(testScript)
	assert.NoError(suite.T(), err)

	actual, err := suite.client.Script.Get(name)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), testScript, actual)

	// Update
	err = suite.client.Script.Update(testScript)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.Script.Delete(name)
	assert.NoError(suite.T(), err)
}

func (suite *NexusClientSuite) TestScriptWithParams() {
	_, err := suite.client.Script.Get("fake-script")
	assert.Error(suite.T(), err)
}

func (suite *NexusClientSuite) TestScriptRunError() {
	name := "test-script"
	testScript := nexus.Script{
		Name: name,
		Content: `
			import groovy.json.JsonSlurper;
			def request = new JsonSlurper().parseText(args);
			assert request.repoName: 'repoName parameter is required';
		`,
		Type: "groovy",
	}

	params := map[string]string{}

	// Create
	err := suite.client.Script.Create(testScript)
	assert.NoError(suite.T(), err)

	// Run
	_, err = suite.client.Script.Run(name, params)
	assert.Error(suite.T(), err)

	params["repoName"] = "test"

	_, err = suite.client.Script.Run(name, params)
	assert.NoError(suite.T(), err)

	// cleanup
	err = suite.client.Script.Delete(name)
	assert.NoError(suite.T(), err)
}
