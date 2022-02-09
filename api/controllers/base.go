package controllers

import (
	"log"
	"net/http"

	"github.com/carloshdurante/geolocation_api/api/database"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) Initialize() {
	database.StartDb()
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
