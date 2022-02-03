package controllers

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
	"github.com/gorilla/mux"
	"github.com/carloshdurante/geolocation_api/api/database"
)

type Server struct {
	Router *mux.Router
	DB *gorm.DB
}

func (server *Server) Initialize(){
	database.StartDb()
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Application's up.")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
