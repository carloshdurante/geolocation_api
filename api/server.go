package api

import (
	"github.com/carloshdurante/geolocation_api/api/controllers"
)

var server = controllers.Server{}

func Run() {
	server.Initialize()
	server.Run(":8080")
}

