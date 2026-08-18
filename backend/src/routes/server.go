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
	s.server.HandleFunc("/api/ping", s.pingHandler)
	s.server.HandleFunc("/api/login/{group}", s.login)
	s.server.HandleFunc("/api/tasks/make", s.makeTask) //maybe also use to update?
	s.server.HandleFunc("/api/tasks/list", s.getTasks)
	s.server.HandleFunc("/api/tasks/random", s.getRandomTask)
	s.server.HandleFunc("/api/tasks/complete", s.completeTask)
	s.server.HandleFunc("/api/tasks/get", s.getTask)
}
