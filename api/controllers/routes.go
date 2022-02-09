package controllers

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/health_check", server.HealthCheck).Methods("GET")
	server.Router.HandleFunc("/addresses", server.GetAllAddress).Methods("GET")
	server.Router.HandleFunc("/addresses", server.CreateAddress).Methods("POST")
	server.Router.HandleFunc("/addresses/{id}", server.DeleteAddress).Methods("DELETE")
	server.Router.HandleFunc("/addresses/{id}", server.GetAddressByID).Methods("GET")
	server.Router.HandleFunc("/addresses/{id}", server.UpdateAddress).Methods("PUT")
}
