package tests

import (
	"net/http"
	"testing"

	"github.com/F0rzend/simple-go-webserver/app/common"
	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/assert"
)

type ErrorChecker = func(assert.TestingT, error, ...any) bool

func HTTPExpect(t *testing.T, handler http.HandlerFunc) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		RequestFactory: newRequestFactoryWithTestLogger(t),
		Reporter:       httpexpect.NewAssertReporter(t),
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
		},
	})
}

type tHelper interface {
	Helper()
}

func AssertErrorFlag(flag common.Flag) ErrorChecker {
	return func(t assert.TestingT, err error, _ ...any) bool {
		if h, ok := t.(tHelper); ok {
			h.Helper()
		}

		if !assert.Error(t, err) {
			return false
		}

		var flagged interface {
			Flag() common.Flag
		}
		assert.ErrorAs(t, err, &flagged)
		return assert.Equal(t, flag, flagged.Flag())
	}
}
