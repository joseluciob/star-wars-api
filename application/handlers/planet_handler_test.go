package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"star-wars-api/application/handlers/middleware"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func TestPlanetIdentifierFail(t *testing.T) {
	r := gin.Default()
	r.GET("/planets/:id", middleware.IdentifierMiddleware())
	req, err := http.NewRequest(http.MethodGet, "/planets/1ffff", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	validationErr := make(map[string]string)

	err = json.Unmarshal(rr.Body.Bytes(), &validationErr)

	assert.Equal(t, rr.Code, http.StatusInternalServerError)
}
