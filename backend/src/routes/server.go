package routes

import (
	"log"
	"net/http"
)

type Server struct {
	*http.ServeMux
}

func NewServer() Server {
	log.Println("Creating new server")
	server := http.NewServeMux()

	return Server{server}
}

func (s Server) StartServer(addr string) error {
	log.Println("Starting server...")
	s.loadRoutes()

	return http.ListenAndServe(addr, s)
}

func (s Server) loadRoutes() {
	log.Println("Loading routes...")
	s.HandleFunc("/ping", s.pingHandler)
	s.HandleFunc("/user/{userId}", s.login)
}
