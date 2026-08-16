package routes

import (
	"log"
	"net/http"
	"noahsTesting/backend/src/database"
)

type Server struct {
	server *http.ServeMux
	db     database.Database
}

func NewServer() Server {
	log.Println("Creating new server")

	return Server{
		server: http.NewServeMux(),
		db:     database.New(),
	}
}

func (s Server) StartServer(addr string) error {
	log.Println("Starting server...")
	s.loadRoutes()

	return http.ListenAndServe(addr, s.server)
}

func (s Server) loadRoutes() {
	log.Println("Loading routes...")
	s.server.HandleFunc("/ping", s.pingHandler)
	s.server.HandleFunc("/api/{group}/login", s.login)
	s.server.HandleFunc("/tasks/make", s.makeTask) //maybe also use to update?
	s.server.HandleFunc("/tasks/list", s.getTasks)
	s.server.HandleFunc("/tasks/random", s.getRandomTask)
	s.server.HandleFunc("/tasks/complete", s.completeTask)
}
