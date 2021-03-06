package handler

import (
	"net/http"
	"testing"

	"net/http/httptest"

	"github.com/at-vudang95/go-food-market-api/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandlerNewErrorHandler(t *testing.T) {
	eh := NewHTTPErrorHandler(infrastructure.NewLogger().Log)
	assert.NotEmpty(t, eh)
}

func TestErrorHandlerStatusNotFound(t *testing.T) {
	eh := NewHTTPErrorHandler(infrastructure.NewLogger().Log)

	// handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		eh.StatusNotFound(w, r)
		return
	})

	// response writer
	rec := httptest.NewRecorder()
	// new request
	req, _ := http.NewRequest("GET", "/", nil)
	// request
	handler(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestErrorHandlerStatusMethodNotAllowed(t *testing.T) {
	eh := NewHTTPErrorHandler(infrastructure.NewLogger().Log)

	// handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		eh.StatusMethodNotAllowed(w, r)
		return
	})

	// response writer
	rec := httptest.NewRecorder()
	// new request
	req, _ := http.NewRequest("GET", "/", nil)
	// request
	handler(rec, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
}
