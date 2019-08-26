package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/checkaayush/sampark/handler"
)

func TestHealth(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/v1/health", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	h := &handler.Handler{getMockDB()}

	// Assertions
	if assert.NoError(t, h.Health(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "OK", rec.Body.String())
	}
}
