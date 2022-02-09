package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/carloshdurante/geolocation_api/api/database"
	"github.com/carloshdurante/geolocation_api/api/models"
	"github.com/carloshdurante/geolocation_api/api/repositories"
	"github.com/gorilla/mux"
)

func (server *Server) GetAllAddress(w http.ResponseWriter, r *http.Request) {
	repository := repositories.AddressRepositoryDb{Db: database.GetDb()}
	addresses, err := repository.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, addresses)
}

func (server *Server) GetAddressByID(w http.ResponseWriter, r *http.Request) {
	repository := repositories.AddressRepositoryDb{Db: database.GetDb()}
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	address, err := repository.FindByID(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, address)
}

func (server *Server) CreateAddress(w http.ResponseWriter, r *http.Request) {
	repository := repositories.AddressRepositoryDb{Db: database.GetDb()}
	address, err := repository.Create(&models.Address{
		Address:    r.FormValue("address"),
		PostalCode: r.FormValue("postal_code"),
		Latitude:   r.FormValue("latitude"),
		Longitude:  r.FormValue("longitude"),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, address)
}

func (server *Server) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	repository := repositories.AddressRepositoryDb{Db: database.GetDb()}
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	address, err := repository.Delete(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, address)
}

func (server *Server) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	repository := repositories.AddressRepositoryDb{Db: database.GetDb()}
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	address, err := repository.Update(id, &models.Address{
		ID:         id,
		Address:    r.FormValue("address"),
		PostalCode: r.FormValue("postal_code"),
		Latitude:   r.FormValue("latitude"),
		Longitude:  r.FormValue("longitude"),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, address)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
