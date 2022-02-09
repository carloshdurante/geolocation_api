package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carloshdurante/geolocation_api/api/controllers"
	"github.com/stretchr/testify/assert"
)

var server = controllers.Server{}

func TestHealthCheckStatusOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/health_check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	server.Initialize()
	server.Router.HandleFunc("/health_check", server.HealthCheck).Methods("GET")
	server.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestHealthCheckNotStatusOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/health_check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	server.Initialize()
	server.Router.HandleFunc("/health_check", server.HealthCheck).Methods("GET")
	http.ResponseWriter(rr).WriteHeader(http.StatusInternalServerError)
	server.Router.ServeHTTP(rr, req)

	assert.NotEqual(t, http.StatusOK, rr.Code)
}
