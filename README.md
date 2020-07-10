# nexus-go-client
>A minimal Golang client for interacting with the Nexus Repository APIs.

[![Build Status](https://travis-ci.org/kaizendorks/nexus-go-client.svg?branch=master)](https://travis-ci.org/kaizendorks/vuecli-in-docker)

## Description

As the name suggests this is a minimal Nexus API client written in go for interacting with the APIs defined here: https://help.sonatype.com/repomanager3/rest-and-integration-api

It does not implement any extra logic other than the one supported by the official Nexus API (its more like a proxy that just passes things back and forward), leaving it up to the users to handle advanced logic. This reduces the amount of code that is not under the user's control, in addition to making this library simpler and easier to keep it in sync with the Nexus APIs.

For the reasons mentioned above and the fact that the Nexus APIs seem to still be in their infancy but evolving at a very fast pace, we have decided to not do anything fancy for re-using code and instead created an API object with matching names (even tho this meant having some long fuction names) to the Nexus API groups. The API objects can be accessed via the nexus client, and in turn provide access to the group APIs E.g:
	Assets
	...
	Repositories
	RepositoryManagement
	...
	SecurityManagement
	SecurityManagementRoles
	SecurityManagementUsers
	Status
	...

In the future we might abstact some of the common fuctionality in addition to creating common interfaces for things like Repository management.

## Usage

```go
import "github.com/kaizendorks/nexus-go-client/nexus"
```

Construct a new Nexus client, then use the various services on the client to access different parts of the Nexus API. For example:

```go
client := nexus.NewClient(nexus.ClientConfig{
	Host:     nexusHost,
	Username: nexusUsername,
	Password: nexusPassword,
})

err := client.Status.StatusWritable()
if err != nil {
	// Report Nexus server is not readable/writable.
}
```

The services of a client correspond to the structure of the Nexus API documentation at: https://help.sonatype.com/repomanager3/rest-and-integration-api

## Contributing

#### Dev environment

We've put together a dockerfile and compose file for starting a dev container in less than a minute.

1. Start dev container: `docker-compose run --rm client`
1. Clean up: `docker-compode down`

#### Running tests

1. Apply style guides: `go fmt ./...`
1. Static analysis: `go vet ./...`
1. Run tests: `go test ./... -v -coverprofile cover.out`
1. Examine code coverage: `go tool cover -func=cover.out`

#### Writing tests

Because we don't have any control over the Nexus API changes and to avoid having to create lots of mock data that will become hard to maintain, we have decided to give preference to Integration tests that directly talk to a Nexus server. These integration tests should be created under the NexusClientSuite.

```go
func (suite *NexusClientSuite) TestStatusReadable() {
	assert.NoError(suite.T(), suite.client.Status.Status())
}
```

In some cases it's not possible or very hard to test certain APIs (e.g one that requires the use of third party tools, like the S3 blobl store), so we also have a second test suite called MockedClientSuite, as the name sugests tests in this suite make use of httptest to mock API responses.

```go
func (suite *MockedClientSuite) TestStatusReadableError() {
	// generate a test server so we can capture and inspect the request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(503)
	}))
	defer ts.Close()

	mockedClient := nexus.NewClient(nexus.ClientConfig{
		Host: ts.URL,
	})

	assert.Error(suite.T(), mockedClient.Status.Status())
}
```
