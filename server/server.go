package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// We create MyServer struct in order to add the run method to it
type MyServer struct {
	server *http.Server
}

// Set up server configuration
// It receives as a parameter the multiplexer (In this case from chi).
func GetServer(mux *chi.Mux) *MyServer {
	s := &http.Server{
		Addr:           ":8000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &MyServer{s}
}

// Method of MyServer that runs the server
func (s *MyServer) Run() {
	log.Fatal(s.server.ListenAndServe())
}
