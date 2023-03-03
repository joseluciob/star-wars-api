package httputil

import (
	err "errors"
	"net/http"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestErrorHandler(t *testing.T) {
	error_test := err.New("test error")

	err := ErrorHandler(error_test)

	assert.Equal(t, err.Code, http.StatusInternalServerError)
}
