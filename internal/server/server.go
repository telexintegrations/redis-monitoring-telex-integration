package server

import (

	"github.com/gorilla/mux" // Using Gorilla Mux for better routing
)

// Server struct holds router instance
type Server struct {
	Router *mux.Router
}

// NewServer initializes and returns a new server instance
func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}

	// Register handlers
	s.routes()

	return s
}

// routes registers all endpoints
func (s *Server) routes() {
	s.Router.HandleFunc("/integration.json", IntegrationHandler).Methods("GET")
	s.Router.HandleFunc("/tick", TickHandler).Methods("POST")
}
