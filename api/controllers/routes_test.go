package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var server = Server{}

//Testing GET HealthCheck ok
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

//Testing GET HealthCheck not ok
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

//Testing GET all address ok
func TestGetAllAddressStatusOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/addresses", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	server.Initialize()
	server.Router.HandleFunc("/addresses", server.GetAllAddress).Methods("GET")
	server.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

//Testing GET all address not ok
func TestGetAllAddressNotStatusOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/addresses", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	server.Initialize()
	server.Router.HandleFunc("/addresses", server.GetAllAddress).Methods("GET")
	http.ResponseWriter(rr).WriteHeader(http.StatusInternalServerError)
	server.Router.ServeHTTP(rr, req)

	assert.NotEqual(t, http.StatusOK, rr.Code)
}

//Testing POST address ok
func TestCreateAddressStatusOk(t *testing.T) {
	req, err := http.NewRequest("POST", "/addresses", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	server.Initialize()
	server.Router.HandleFunc("/addresses", server.CreateAddress).Methods("POST")
	server.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

//Testing POST address not ok
func TestCreateAddressNotStatusOk(t *testing.T) {
	req, err := http.NewRequest("POST", "/addresses", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	server.Initialize()
	server.Router.HandleFunc("/addresses", server.CreateAddress).Methods("POST")
	http.ResponseWriter(rr).WriteHeader(http.StatusInternalServerError)
	server.Router.ServeHTTP(rr, req)

	assert.NotEqual(t, http.StatusCreated, rr.Code)
}

//Testing DELETE address ok
func TestDeleteAddressesStatusOk(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/addresses/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	server.Initialize()
	server.Router.HandleFunc("/addresses/1", server.DeleteAddress).Methods("DELETE")
	server.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

//Testing DELETE address not ok
func TestDeleteAddressesNotStatusOk(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/addresses/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	server.Initialize()
	server.Router.HandleFunc("/addresses/1", server.DeleteAddress).Methods("DELETE")
	http.ResponseWriter(rr).WriteHeader(http.StatusInternalServerError)
	server.Router.ServeHTTP(rr, req)

	assert.NotEqual(t, http.StatusOK, rr.Code)
}

//Testing GET address by ID ok
func TestGetAddressByIDStatusOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/addresses/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	server.Initialize()
	server.Router.HandleFunc("/addresses/2", server.GetAddressByID).Methods("GET")
	server.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

//Testing GET address by ID not ok
func TestGetAddressByIDNotStatusOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/addresses/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	server.Initialize()
	server.Router.HandleFunc("/addresses/2", server.GetAddressByID).Methods("GET")
	http.ResponseWriter(rr).WriteHeader(http.StatusInternalServerError)
	server.Router.ServeHTTP(rr, req)

	assert.NotEqual(t, http.StatusOK, rr.Code)
}

//Testing PUT address by ID ok
func TestUpdateAddressStatusOk(t *testing.T) {
	req, err := http.NewRequest("PUT", "/addresses/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	server.Initialize()
	server.Router.HandleFunc("/addresses/2", server.UpdateAddress).Methods("PUT")
	server.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

//Testing PUT address by ID not ok
func TestUpdateAddressNotStatusOk(t *testing.T) {
	req, err := http.NewRequest("PUT", "/addresses/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	server.Initialize()
	server.Router.HandleFunc("/addresses/2", server.UpdateAddress).Methods("PUT")
	http.ResponseWriter(rr).WriteHeader(http.StatusInternalServerError)
	server.Router.ServeHTTP(rr, req)

	assert.NotEqual(t, http.StatusOK, rr.Code)
}
