package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayerErrorMarshaling(t *testing.T) {
	t.Parallel()

	const (
		errorMessage = "error message"
		statusCode   = http.StatusInternalServerError
	)

	expectJSON := []byte(fmt.Sprintf(`{"error":%q}`, errorMessage))

	appError := NewApplicationError(statusCode, errorMessage)
	actual, err := json.Marshal(appError)

	assert.NoError(t, err)
	assert.Equal(t, expectJSON, actual)
}
