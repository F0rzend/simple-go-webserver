package tests

import (
	"net/http"
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gavv/httpexpect"
)

func HTTPExpect(t *testing.T, handler http.HandlerFunc) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		Reporter: httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
		},
	})
}

func ExpectApplicationError(t *testing.T, expectedStatus int, err error) {
	t.Helper()

	if expectedStatus == 0 {
		assert.NoError(t, err)
	} else {
		require.IsType(t, common.ApplicationError{}, err)
		err, ok := err.(common.ApplicationError)
		if !ok {
			t.Fatalf(
				"Object expected to be of type %T, but was %T",
				common.ApplicationError{},
				err,
			)
		}
		assert.Equal(t, expectedStatus, err.HTTPStatus)
	}
}
