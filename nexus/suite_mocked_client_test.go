package nexus_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MockedClientSuite struct {
	suite.Suite
}

func TestMockedClientSuite(t *testing.T) {
	suite.Run(t, new(MockedClientSuite))
}
