package tests

import (
	"context"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	e             func() *httpexpect.Expect
	tearDownSuite func()
}

func (s *TestSuite) SetupSuite() {
	ctx := context.Background()

	app, err := setupTestEnvironment(ctx)
	require.NoError(s.T(), err)

	s.e = func() *httpexpect.Expect {
		return httpexpect.New(s.T(), app.URI)
	}
	s.tearDownSuite = func() {
		app.close(ctx)
	}
}

func (s *TestSuite) TearDownSuite() {
	s.tearDownSuite()
}

func TestApplication(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
