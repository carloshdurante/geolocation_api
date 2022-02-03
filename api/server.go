package api

import (
	"os"

	"github.com/carloshdurante/geolocation_api/api/controllers"
)

var server = controllers.Server{}

func Run() {
	server.Initialize()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}
