package controllers

import (
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/carloshdurante/geolocation_api/api/database"
	"github.com/carloshdurante/geolocation_api/api/models"
	"github.com/carloshdurante/geolocation_api/api/repositories"
	"github.com/carloshdurante/geolocation_api/api/services/address"
	enrichmentAddress "github.com/carloshdurante/geolocation_api/api/services/enrichment_address"
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
	var newAddress models.Address
	json.NewDecoder(r.Body).Decode(&newAddress)

	repository := repositories.AddressRepositoryDb{Db: database.GetDb()}
	address, err := address.Create(&newAddress, &repository)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	db := enrichmentAddress.DbConnection{Db: database.GetDb()}
	_, err = db.StartEnrichment(address.Address, address.ID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response_address := map[string]string{
		"id": strconv.FormatUint(address.ID, 10),
	}
	respondWithJSON(w, http.StatusCreated, response_address)
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
	var newAddress models.Address
	json.NewDecoder(r.Body).Decode(&newAddress)

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	db := enrichmentAddress.DbConnection{Db: database.GetDb()}
	address, err := db.StartEnrichment(newAddress.Address, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response_address := map[string]uint64{
	"id": address.ID,
	}

	respondWithJSON(w, http.StatusOK, response_address)
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
