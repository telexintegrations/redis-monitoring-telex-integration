package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/BerylCAtieno/redis-monitor/internal/monitor"
)

// Server represents the HTTP server
type Server struct {
	RedisMonitor *monitor.RedisMonitor
	Alert        *monitor.Alert
}

// NewServer creates a new instance of the server
func NewServer(redisMonitor *monitor.RedisMonitor, alert *monitor.Alert) *Server {
	return &Server{
		RedisMonitor: redisMonitor,
		Alert:        alert,
	}
}

// Start initializes the server and starts listening
func (s *Server) Start(port int) {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/tick", TickHandler).Methods("GET")
	r.HandleFunc("/alert", AlertHandler(s.RedisMonitor, s.Alert)).Methods("POST")

	// Start the server
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Server listening on %s...", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
