package main


import (
	"log"
	"net/http"
	"github.com/BerylCAtieno/redis-monitor/internal/server"
)

func main() {
	// Start the server
	srv := server.NewServer()

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", srv.Router))
}
