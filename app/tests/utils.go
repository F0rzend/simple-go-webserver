package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertStatus(
	t *testing.T,
	handler http.Handler,
	method string,
	path string,
	body any,
	expected int,
) (http.ResponseWriter, *http.Request) {
	t.Helper()

	requestBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRequest(method, path, bytes.NewReader(requestBody))
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	actual := w.Code

	if actual != expected {
		errorMessage := fmt.Sprintf("Expected HTTP status code %d but received %d", expected, actual)
		errorMessage += fmt.Sprintf("\n\nURL: %s", r.URL.String())
		if requestBody != nil {
			errorMessage += fmt.Sprintf("\nRequest: %#v", body)
		}
		if w.Body.Len() > 0 {
			errorMessage += fmt.Sprintf("\nResponse: %s", w.Body.String())
		}
		assert.Fail(t, errorMessage)
	}

	return w, r
}

func must[value any](val value, err error) value {
	if err != nil {
		panic(err)
	}
	return val
}
