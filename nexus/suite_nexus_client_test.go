package nexus_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/kaizendorks/nexus-go-client/nexus"
)

const DefaultDockerHost = "unix:///var/run/docker.sock"

type NexusClientSuite struct {
	suite.Suite
	client nexus.Client
}

func (suite *NexusClientSuite) waitForNexusServer() {
	tries := 0
	err := suite.client.Status.StatusWritable()
	for tries < 5 && err != nil {
		fmt.Println("Nexus server not ready, sleeping 15 seconds.")
		time.Sleep(15 * time.Second)
		err = suite.client.Status.StatusWritable()
		tries++
	}

	if err != nil {
		panic(err)
	}
}

func (suite *NexusClientSuite) SetupSuite() {
	fmt.Printf("\033[1;36m%s\033[0m", "> Start NexusClientSuite setup\n")
	suite.client = nexus.NewClient(nexus.ClientConfig{
		Username: "admin",
		Password: "admin123",
		Host:     "http://nexus:8081",
	})

	suite.waitForNexusServer()
	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func TestNexusClientSuite(t *testing.T) {
	suite.Run(t, new(NexusClientSuite))
}
