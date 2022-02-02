package controllers

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/health_check", server.HealthCheck).Methods("GET")
}