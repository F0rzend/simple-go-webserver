package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertStatus(
	t *testing.T,
	w http.ResponseWriter,
	r *http.Request,
	expectedStatus int,
) {
	t.Helper()

	recorder, ok := w.(*httptest.ResponseRecorder)
	if !ok {
		t.Fatal("writer is not *httptest.ResponseRecorder")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	actual := recorder.Code

	if actual != expectedStatus {
		errorMessage := fmt.Sprintf("Expected HTTP status code %d but received %d", expectedStatus, actual)
		errorMessage += fmt.Sprintf("\n\nURL: %s", r.URL.String())

		if len(body) != 0 {
			errorMessage += fmt.Sprintf("\nRequest: %s", body)
		}
		if recorder.Body.Len() > 0 {
			errorMessage += fmt.Sprintf("\nResponse: %s", recorder.Body.String())
		}
		assert.Fail(t, errorMessage)
	}
}

func PrepareHandlerArgs(
	t *testing.T,
	method string,
	path string,
	body any,
) (*httptest.ResponseRecorder, *http.Request) {
	t.Helper()

	requestBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRequest(method, path, bytes.NewReader(requestBody))
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	return w, r
}
